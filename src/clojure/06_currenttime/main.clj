(import '[java.util Calendar]
        '[java.text SimpleDateFormat])

(defn -get-time
  []
  (.format (SimpleDateFormat. "yyyyMMdd_HHmmss")
           (.getTime (Calendar/getInstance))))

(let [args *command-line-args*]
  (if (< (count args) 1)
    (println "need 1 arguments")
    (spit (str (-get-time) ".txt")
          (first args))))
