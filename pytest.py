import sys,os
from dotenv import load_dotenv

version_tag = str(sys.argv[1])

with open(".env", "w") as f:
    f.write("version={}".format(version_tag))
    
load_dotenv()

tag_version = os.getenv('version')
print(tag_version+" Passed!!")
