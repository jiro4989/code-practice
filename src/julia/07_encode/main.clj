(spit "out_utf8.txt"
      (slurp "in_cp932.txt" :encoding "Shift_JIS")
      :encoding "UTF-8")

(spit "out_cp932.txt"
      (slurp "in_utf8.txt" :encoding "UTF-8")
      :encoding "Shift_JIS")
