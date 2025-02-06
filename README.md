# go-led
LED Scrolling Banner for GO

Written to continue my exploration and learning of the GO language. Please note with this repo, I am EXTREMELY NEW to GO.  While this demo does work, and seems to work well, it's written by a GO novice, and as and when my GO skills improve i'll refactor this.

This is a console mode application.  For controlling console i/o I use https://github.com/rivo/tview/wiki 

The idea here is to display a string of text, scrolling (initially) from right-left in a simulated LED scrolling display using coloured characters to simulate LEDs being lit and LEDs that are not lit.

This may be enhanced to allow piping text to it from the command line;  as well as other controls for speed, direction, colour and size.

A 'small' character set is added in _characters.go_ A-Z and 0-9. Only capital letters, numerals and a space character.  This can be extended easily to add more characters, but if i'm honest, I got a bit bored drawing these out :) 

![Sample Output](./go-led.gif)

As can be seen by the above demo, there remains a problem to resolve with this - that is, when the text reaches the left edge, it doesn't scroll off the left edge but indeed resets to the start of the right edge again.

While no visible instructions are given, the [CursorUp] & [CursorDown] keys can be used to change the speed of the text scrolling.