import codecs
import json
import os
import sys
import soundfile
from random import *

for line in open('/data/gow/clean_deepspeech/DeepSpeech/es_ar_male/line_index.tsv', 'r'):
	line = line.decode('utf-8')
	segments = line.strip().split()
	text = ' '.join(segments[1:]).lower()
        for letter in ',?-!.:':
            text=text.replace(letter,'')
	audio_filepath = '/data/gow/clean_deepspeech/DeepSpeech/es_ar_male'+'/'+segments[0]+'.wav'
	#print audio_filepath
	audio_data,samplerate = soundfile.read(audio_filepath)
	duration =float(len(audio_data)) / samplerate
	linevali = json.dumps({
                            'audio_filepath': audio_filepath,
                            'duration': duration,
                            'text': text
                             },ensure_ascii=False)
	x = randint(0,1000000)
        if (x%20==19 or x%20==12):
            with codecs.open('/data/gow/clean_deepspeech/DeepSpeech/manifest.test','a+',encoding ='utf-8') as out_file:
                out_file.write(linevali+'\n')
        elif (x%20==8 or x%20==5):
            with codecs.open('/data/gow/clean_deepspeech/DeepSpeech/manifest.vali','a+',encoding ='utf-8') as out_file:
                out_file.write(linevali+'\n')
        else :
            with codecs.open('/data/gow/clean_deepspeech/DeepSpeech/manifest.train', 'a+',encoding ='utf-8') as out_file:
                out_file.write(linevali+ '\n')


	
