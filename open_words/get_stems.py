"""

get_stems

Small helper function to get a list of unique stems from the inflects data

"""

import re

# from "./inflects" import Inflects # change: just directly importing inflects
# from inflects import Inflects
from inflects import Inflects

def get_stems(word, inflects=Inflects):

	stems = [ word ]

	# Sort by length of ending
	inflects.sort(key=lambda x: len(x['ending']))

	if not inflects:
		print("Error importing inflects")

	for inflect in inflects:
		if word.endswith( inflect['ending'] ):
			stem = re.sub( inflect['ending'] + "$", "", word )
			if word != stem and stem not in stems:
				print(stem, stems)
				stems.append( stem )

	return stems

print(get_stems('habeo'))