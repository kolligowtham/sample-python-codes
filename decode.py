import json
import codecs
import os
manifest_path = '/data/gow/clean_deepspeech/DeepSpeech/manifest.train'
#manifest = []
for json_line in codecs.open(manifest_path, 'r', 'utf-8'):
	#print type(json_line)
	#json_line=json_line.encode('latin-1')
	#print json_line
	try:
		print json_line
		json_data = json.loads(json_line)
		print ('try')
	except Exception as e:
		raise IOError("Error reading manifest: %s" % str(e))
    #manifest.append(json_data)
#manifest_jsons=manifest
#for line_json in manifest_jsons:
    
