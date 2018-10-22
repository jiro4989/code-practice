(let [args *command-line-args*]
  (if (< (count args) 2)
    (println "need 2 arguments.")
    (println (reduce + (map #(Integer. %) args)))))

; マクロを使って少し読みやすく
(let [args *command-line-args*]
  (if (< (count args) 2)
    (println "need 2 arguments.")
    (->> args
        (map #(Integer. %))
        (reduce +)
        println
        )))
