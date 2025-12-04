defmodule Solve do
  def find_max(lst, l_max, r_max, maxfd) do
    if lst == [] do
      IO.inspect(r_max * 10 + l_max)
      r_max * 10 + l_max
    else
      cur = Enum.at(lst, 0)
      if cur >= r_max do
	# l_max = max l_max, r_max
	# r_max = cur
	find_max(List.delete_at(lst, 0), maxfd, cur, max(cur,maxfd))
      else # cur <= r_max
	# l_max = l_max
	# r_max = r_max
	find_max(List.delete_at(lst, 0), l_max, r_max, max(cur, maxfd))
      end
    end
  end
  def max_joltage(bank) do
    #IO.inspect(bank)
    bank = Enum.reverse(bank)
    find_max(List.delete_at(bank, 0), Enum.at(bank, 0), 0, Enum.at(bank, 0))
  end
end

File.read!("input.txt")
|> String.split()
|> Enum.map(& String.to_integer(&1) |> Integer.digits() |> Solve.max_joltage())
|> Enum.sum()
|> IO.puts()
