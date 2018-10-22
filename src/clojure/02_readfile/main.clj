(require '[clojure.string :as str])

(let [args *command-line-args*] 
  (if (< (count args) 1)
    (println "need 1 arguments.")
    (doseq [[i line] (map-indexed vector (-> args
                                             first
                                             slurp
                                             (str/split #"\n")))]
      (printf "%03d:%s\n" i line)
      )))
