# Compilation Steps
Testing was done on Windows 11
## Arduino
1. Install the Arduino IDE https://www.arduino.cc/en/software
	1. My Version is 2.2.1
2. On an Arduino Uno connect pin 7 to the output(middle) pin of the DHT22 sensor then connect the 5 volt power pin to the positive(left) pin and the GND pin to the negative(right) pin.
3. Connect the board to your computer using a USB cable
4. Then in the IDE open the file in this directory arduino/SensorTest called SensorTest.ino
5. Verify that the IDE is pointing to the connected board in the dropdown menu in the top left
6. Click the arrow going to the right to upload the file to the board
7. After that is finished you can open Serial Monitor which is located in the top right to see the output from the board.
## Web
1. Install Golang on your system
	1. https://go.dev/dl/
	2. My Go version is go1.21.6 windows/amd64
2. Inside of a terminal navigate to the web directory and run the command 'go run main.go'
3. This will start the web server on local host port 8000
	1. http://localhost:8000/
4. Use Control + C to stop the process(Or Command + Shift + C on mac)
# Output
## Arduino
![[Pasted image 20240219141851.png]]
## Web
![[Screenshot 2024-02-19 142050.png]]
![[Pasted image 20240219142306.png]]