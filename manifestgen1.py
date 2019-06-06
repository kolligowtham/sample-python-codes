import codecs
import json
import os
import sys
import soundfile
from random import *
data_dir ='/data/gow/clean_deepspeech/DeepSpeech/es_ar_female'
for subfolder, _,filelist in sorted(os.walk(data_dir)):
    text_filelist = [
        filename for filename in filelist if filename.endswith('.tsv')
    ]
    if len(text_filelist) > 0:
        i=0 
        for name_of_file in text_filelist :
            text_filepath = os.path.join(data_dir,subfolder,text_filelist[i])
            i=i+1
            
	    print text_filepath
            for line in open(text_filepath):
                try:
                    line = line.decode('utf-8')
                    segments = line.strip().split('\t')
                    text = ' '.join(segments[1:]).lower()
                    for letter in ',?:-!.':
                        text = text.replace(letter,'')
                    audio_filepath = os.path.join(data_dir,subfolder,segments[0]+'.wav')
                    audio_data,samplerate = soundfile.read(audio_filepath)
                    duration = float(len(audio_data)) / samplerate
                    linevali=json.dumps({
                            'audio_filepath': audio_filepath,
                            'duration': duration,
                            'text': text
                             },ensure_ascii=False)
                    x=randint(0,1000000)
                    if (x%20==4 or x%20==13):
                        with codecs.open('/data/gow/clean_deepspeech/DeepSpeech/manifest.test','a+',encoding='utf-8') as out_file:
                            out_file.write(linevali+'\n')
                    elif (x%20==1 or x%20==6):
                        with codecs.open('/data/gow/clean_deepspeech/DeepSpeech/manifest.vali','a+',encoding='utf-8') as out_file:
                            out_file.write(linevali+'\n')
                    else : 
                        with codecs.open('/data/gow/clean_deepspeech/DeepSpeech/manifest.train', 'a+',encoding='utf-8') as out_file:
                            out_file.write(linevali+ '\n')


                except Exception as e:
                    print e
                    print 'segments =',segments

