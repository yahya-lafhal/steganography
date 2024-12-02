# Steganography

Steganography is the practice of concealing information within another message or physical object to avoid detection.

This program does that with a message (string) as the information we want to hide, and we will hide it inside an image (png).

It offers bother encoding and decoding.

# Running the program

Assuming you have go installed, open your terminal then execute the following command in the project folder

```bash
$ go run .
```

You can also build the program to have a .exe that you can use

```bash
$ go build .
```

# Encoding 

Once you run the program, you will have to chose to encode, then you will enter the file name of the image you want to encode into. You will then be asked to enter the message. It will create another image file with the same name but with the "encoded_" prefix and in the png format. The two images are indistinguishable by the human eye.

# Decoding 

Once you run the program, you will have to chose to decode, then you will enter the file name of the image you want to decode. It will look for a message written in the same format this program encodes.

# TODO

- Handle more image formats (jpeg, jpg, ...)
- Add an encryption alogrithm 
- Clean and optimize code