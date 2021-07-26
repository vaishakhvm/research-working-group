import sys

version_tag = str(sys.argv[1])

with open(".env", "w") as f:
    f.write("version={}".format(version_tag))
