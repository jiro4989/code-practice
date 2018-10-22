(require '[clojure.string :as str])

(def outfile "out.txt")

; ファイルを空にする
(spit outfile "")

(let [args *command-line-args*]
  (if (< (count args) 1)
    (println "need 1 arguments.")
    (doseq [[i line] (map-indexed vector (-> args
                                             first
                                             slurp
                                             (str/split #"\n")))]
      (spit outfile (format "%03d:%s\n" i line) :append true :encoding "UTF-8")
      )))
