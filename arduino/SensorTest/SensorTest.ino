#include <DHT.h>

#define DHTPIN 7
#define DHTTYPE DHT22

DHT dht(DHTPIN, DHTTYPE);

float temp;
float hum;

void  setup()
{
  Serial.begin(9600);
  dht.begin();
 
}

void loop()
{
  Serial.println();

  temp = dht.readTemperature();
  hum = dht.readHumidity();

  Serial.print("Humidity (%): ");
  Serial.println(hum, 2);

  Serial.print("Temperature  (C): ");
  Serial.println(temp, 2);

  delay(2000);

}