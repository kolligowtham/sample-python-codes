import sox
import os 
data_dir = '/data/gow/clean_deepspeech/DeepSpeech/'
for (sub_folder, _,filelist) in sorted(os.walk(data_dir)):
	for filename in filelist:
		if(filename.endswith('.wav')):
			newname = filename.replace('sr16'.'')