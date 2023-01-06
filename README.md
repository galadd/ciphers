# <center> Ciphers </center>
This console application allows users to easily encrypt and decrypt messages using a variety of different ciphers. </br>
Includes [the following ciphers](https://github.com/galadd/ciphers/tree/main/ciphers).


## Running the application
To run the console application, follow these steps:

1. Open a terminal or command prompt.

2. Navigate to the directory where the application is located.

3. Run the application 
    * On Windows, run `ciphers.exe`
    * On Linux, run `./ciphers`

4. When the application starts, it will prompt you to enter a message that you want to encrypt or decrypt. Type your message and press Enter.

5. Next, the application will display a list of available ciphers. Choose a cipher by typing its number and pressing Enter.

6. The application will then perform the chosen cipher on the message and display the resulting output.

That's it! You can now experiment with different ciphers and messages to see how they work.

## Testing

To run the tests for the project, use the command:

```
go test ./tests/<cipher>
```
or

```
cd tests
cd <cipher>
go test
```