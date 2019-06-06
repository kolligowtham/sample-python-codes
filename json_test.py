import codecs
import json
import os
import sys
import soundfile
text_path='/data/gow/clean_deepspeech/DeepSpeech/DVDCorpusDimex100/CorpusDimex100/s001/texto/comunes/s00110.txt'
for line in open(text_path):
	print line
	text =line.replace("\r\n",'')
	print text
	text = text.decode('ISO-8859-1')
	for letter in '.,?!-:':
		text =text.replace(letter,'')
	print text
	#text = text.decode('ISO-8859-1')
	text=text.rstrip()
	linevali = json.dumps({'text':text
		},ensure_ascii=False)
	print type(linevali)
	with codecs.open('/data/gow/clean_deepspeech/DeepSpeech/manifest.exp','a+',encoding='utf-8') as out_file:
					out_file.write(linevali+'\n')

	