(require '[clojure.java.io :as io]
         '[clojure.string :as str])

(with-open [fin (io/reader "access_log")]
  (doseq [line (line-seq fin)]
    (-> line
        (str/replace #"^([^\s]+) .*" "$1")
        println)))
