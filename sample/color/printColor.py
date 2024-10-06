from .color import Color

def printRed(text: str) -> None:
   print(Color.letters.red + text + Color.reset)

def printGreen(text: str) -> None:
   print(Color.letters.green + text + Color.reset)

def printBlackOnWhite(text: str) -> None:
   print(Color.backgroud.lightgrey + Color.letters.black + text + Color.reset)
