#!/usr/local/bin/julia

using Printf

if length(ARGS) < 2
  println("need 2 arguments.")
  exit(1)
end

open(ARGS[1], "r") do inf
  open(ARGS[2], "w") do of
    i = 0
    for line in eachline(inf)
      i = i + 1
      s = @sprintf "%03d:%s\n" i line
      write(of, s)
    end
  end
end
