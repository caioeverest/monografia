#include <ArduinoJson.h>
#include <ESP8266WiFi.h>
#include <PubSubClient.h>
#include <SPI.h>

#define WIFI_MAX_CONN_ATTEMPTS 20

//GPIO-NODE-MAPPING
#define D0 16
#define D1 5
#define D2 4
#define D3 0
#define D4 2
#define D5 14
#define D6 12
#define D7 13

// Update these with values suitable for your network.
byte mac[]    = {  0xDE, 0xED, 0xBA, 0xFE, 0xFE, 0xED };

//wifi
char nome_wifi[] = "SSID_da_rede";
char password_wifi[] = "SenhaDaRede";
int connectionAttempts = 0;
//mqtt
const char* mqttServer = "HOST_MQTT";
const int mqttPort = 6883;
const char* mqttUser = "USER_MQTT";
const char* mqttPassword = "PASS_MQTT";
const char* topico = "IDENTIFICADOR_SALA";

const int led = D2;

WiFiClient espClient;
PubSubClient mqttClient(espClient);
StaticJsonBuffer<200> jsonBuffer;

void printWiFiStatus()
{
  if (WiFi.status() != WL_CONNECTED) {
    Serial.println("A conexão wifi falhou :c pois: ");
    Serial.println(WiFi.status());
  } else {
    Serial.println("SSID: ");
    Serial.println(WiFi.SSID());

    Serial.println("IP Local: ");
    IPAddress ip = WiFi.localIP();
    Serial.println(ip);

    Serial.println("Forca do sinal (RSSI): ");
    long rssi = WiFi.RSSI();
    Serial.println(rssi);
    Serial.println(" dBm");
  }
}

void callback(char* topic, byte* payload, unsigned int length) {
  Serial.print("Comando recebido em sala: ");
  Serial.println(topic);

  char data[2000];
  for (int i = 0; i < length; i++) {
    data[i] = (char)payload[i];
  }

  Serial.println(data);

  JsonObject& mensagem = jsonBuffer.parseObject(data);
  if (!mensagem.success()) {
    Serial.println("Parser da mensagem falou...");
    return;
  }
  bool status = mensagem["ativo"];

  Serial.print("Energia foi definida como: ");
  Serial.print(status);

  if(status == true) {
    digitalWrite(led, HIGH);
    Serial.println("Ligando led!");
  } else {
    digitalWrite(led, LOW);
    Serial.println("Desligando led :c");
  }
  
  Serial.println();
  Serial.println("-----------------------");
}

void setup()
{
  Serial.begin(115200);

  Serial.println("Tentando conectar ao wifi: ");
  Serial.println(nome_wifi);

  WiFi.begin(nome_wifi, password_wifi);
  Serial.println("WiFi foi inicializado.");

  while (WiFi.status() != WL_CONNECTED && ++connectionAttempts <= WIFI_MAX_CONN_ATTEMPTS) {
    delay(500);
    Serial.println("Tentando conectar ao WiFi... Tentativa numero:" + connectionAttempts);
  }
  printWiFiStatus();

  delay(1500);
  mqttClient.setServer(mqttServer, mqttPort);
  mqttClient.setCallback(callback);

  while (!mqttClient.connected()) {
    Serial.println("Connecting to MQTT...");

    if (mqttClient.connect("Lampada-Client", mqttUser, mqttPassword )) {
      Serial.println("Conectado com sucesso ao servidor MQTT");
    } else {
      Serial.println("A conxao com o servidor MQTT nao foi estabelecida");
      Serial.print(mqttClient.state());
      delay(2000);
    }
  }

  mqttClient.subscribe(topico);
  pinMode (led, OUTPUT);
}

void loop()
{
  mqttClient.loop();
}
