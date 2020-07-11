from parse import Parse
import sys

# print('reached')
# Sunt geminae Somni portae
# print(sys.argv[1])
parser = Parse()
print(parser.parse_line(sys.argv[1])) #))
sys.stdout.flush()
