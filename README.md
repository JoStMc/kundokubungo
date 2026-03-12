# kundokubungo
訓読文語 – a go tool for going between 訓読文 and 書き下し文

---

## 返り点

#### レ点
Character is parsed after the subsequent character.

#### 一二(三)点
Each number is not parsed until the previous number has been parsed. 

**Note:** Apparently there is at least one instance where there is a sentence which goes up to 九, then 下; this may be a mistake, so in any case where the tenth is 下, it should be replaced with 十. 

Full: 一二三四五六七八九十

#### 上(中)下点
Each is parsed only once the previous mark has been parsed – 中 may be omitted. These are typically used to wrap 一二.

#### 甲乙(丙)点
Each is parsed only once the previous mark has been parsed. Typically they wrap 上(中)下, but they may be swapped with 天地人 since there are only 3 marks in the sequence. The sequence only goes to 己 as is. I will add more if I come across a passage which uses more.

Full: 甲乙丙丁戊己庚辛壬癸

#### 天地(人)点
Each is parsed once the previous mark has been parsed. Typically these wrap 甲乙(丙)点.

#### 元亨(利貞)点
Each is parsed once the previous mark has been parsed. Typically these wrap 天地(人)点.

#### 乾坤点
Each is parsed once the previous mark has been parsed. Typically these wrap  元亨(利貞)点
