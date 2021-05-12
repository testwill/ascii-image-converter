# ascii-image-converter

ascii-image-converter is a command-line tool that converts images into ascii art and prints them out onto the console. It is cross-platform so both Windows and Linux distributions are supported

Currently, the tool supports PNG, JPEG/JPG, WEBP and BMP image formats

<br>

### Example ([Source](https://medium.com/@sean.glancy/practical-applications-of-binary-trees-3097cf663062))
![Example](https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_images/tree.png)

### ASCII Art:
![Example](https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_images/ascii_tree.png)

<br>


## Installation
First download the executables from [here](https://github.com/TheZoraiz/ascii-image-converter/releases/tag/v1.1.1), and follow the steps with respect to your OS.

### Linux
Extract Executables.zip and open the "Linux" directory.

Now, open a terminal in the same directory and execute this command:

```
sudo cp ascii-image-converter /usr/local/bin/
```
Now you can use ascii-image-converter in the terminal. Execute "ascii-image-converter -h" for more details.

### Windows
Extract Executables.zip and open the "Windows" folder. Copy the path to folder from the top of the file explorer and follow these instructions:
* In Search, search for and then select: System (Control Panel)
* Click the Advanced system settings link.
* Click Environment Variables. In the section User Variables find the Path environment variable and select it. Click "Edit".
* In the Edit Environment Variable window, click "New" and then paste the path of the folder that you copied initially.
* Afterwards, you can use it anywhere by typing "ascii-image-converter" in command prompt. Note: Make sure you restart the command prompt.

<br>

## Usage

Tip: Decrease font size or zoom out of terminal for maximum quality ascii art

To convert an image into ascii format, the usage is as follows:
```
ascii-image-converter [path to image]
```
Example
```
ascii-image-converter myImage.jpeg
```
<br>

### Flags

#### --complex OR -c
Print the image with a wider array of ascii characters. Sometimes improves accuracy.
```
ascii-image-converter [path to image] -c
```


#### --dimensions OR -d
Set the width and height of the printed ascii image in character lengths.
```
ascii-image-converter [path to image] -d <width>,<height>
# Or
ascii-image-converter [path to image] --dimensions <width>,<height>
```
Example:
```
ascii-image-converter [path to image] -d 100,30
```

#### --save OR -S
Save the image ascii art in a file ascii-image.txt in the same directory
```
ascii-image-converter [path to image] --save
# Or
ascii-image-converter [path to image] -S
```
<br>

You can combine commands as well
```
ascii-image-converter [path to image] -Scd 100,30
```

<br>

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Packges used

[github.com/spf13/viper](https://github.com/spf13/viper)

[github.com/spf13/cobra](https://github.com/spf13/cobra)

[github.com/mitchellh/go-homedir](https://github.com/mitchellh/go-homedir)

[github.com/nathan-fiscaletti/consolesize-go](https://github.com/nathan-fiscaletti/consolesize-go)

[github.com/nfnt/resize](https://github.com/nfnt/resize)


## License
[Apache-2.0](https://github.com/TheZoraiz/ascii-image-converter/blob/master/LICENSE)