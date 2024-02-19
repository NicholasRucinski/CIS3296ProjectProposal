# Compilation Steps
Testing was done on Windows 11 and Manjaro Linux(Vulcan 23.1.3)
## Arduino
1. Install the Arduino IDE https://www.arduino.cc/en/software
	1. My Version is 2.2.1
2. On an Arduino Uno connect pin 7 to the output(middle) pin of the DHT22 sensor then connect the 5 volt power pin to the positive(left) pin and the GND pin to the negative(right) pin.
3. Connect the board to your computer using a USB cable
4. Then in the IDE open the file in this directory arduino/SensorTest called SensorTest.ino
5. Verify that the IDE is pointing to the connected board in the dropdown menu in the top left
6. Click the arrow going to the right to upload the file to the board
	1. Missing Library
		1. Either point the IDE to the library folder in this repo (File->Preferences and changed sketchbook location to arduino/SensorTest)
		2. Or install the DHT sensor library by Adafruit (Sketch->Include Library->Manage Libraries, or hit the books on the left, and search for "DHT sensor library")![[Images/Pasted image 20240219150026.png]]
	2. Linux device permission denied
		1. Add the current user to either the dialout or uucp group depending on your distrobution and restart
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
![[Images/Pasted image 20240219141851.png]]
## Web
![[Images/Screenshot 2024-02-19 142050.png]]
![[Images/Pasted image 20240219142306.png]]