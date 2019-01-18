import cloudconvert
import os
import sys


log = open("C:\\Users\\Kuba\\IdeaProjects\\project\\logs.txt","w")
log.write("start")
## get input ##
from pip._vendor.distlib.compat import raw_input
from argparse import ArgumentParser

zplString=""
if __name__ == "__main__":
    from codecs import decode
    import sys
    zplString = decode(sys.argv[1], 'unicode_escape')
    log.write(zplString)

filename = "converter\\workspace\\"+zplString
output=zplString.replace(".txt",".doc")
outputfile="converter\\workspace\\"+output
log.write(filename+"\n")
log.write(outputfile+"\n")


## delete only if file exists ##
log.write(os.getcwd()+"\n")
#if os.path.exists(filename):
#    os.remove(filename)
 #   if not os.path.exists(filename):
 #       print("Removing temporary files")

api = cloudconvert.Api('RUAm3o6VYcM8fmtEUMDZCyOShNUmgoZNcagDTHy9vNkjZt21F6BjiqpdShfiGoWU')

log.write("Converting"+"\n")
process = api.convert({
    'inputformat': 'txt',
    'outputformat': 'doc',
    'file': open(filename, 'rb')
})

process.wait()
process.download(outputfile)

if os.path.exists(os.getcwd()+"\\"+outputfile):
    log.write("File converted"+"\n")