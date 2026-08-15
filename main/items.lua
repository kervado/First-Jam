--------------------------------------------------------------------------------
-- items.lua
--
-- Shared definitions for the three carryable item types. Required by the
-- controller, the desks and the player so type ids, colours and carry limits
-- live in exactly one place.
--------------------------------------------------------------------------------

local M = {}

-- item type ids
M.COFFEE = hash("coffee")
M.PAPER  = hash("paper")
M.CAT    = hash("cat")

-- colour used both for a desk's request icon and the player's carry indicator
M.COLORS = {
	[M.COFFEE] = vmath.vector4(0.95, 0.78, 0.30, 1.0),  -- amber
	[M.PAPER]  = vmath.vector4(0.95, 0.95, 0.90, 1.0),  -- white sheet
	[M.CAT]    = vmath.vector4(0.92, 0.55, 0.22, 1.0),  -- ginger
}

-- how many units the assistant carries after touching each source, and the
-- hard carry limit for that item type (they are the same thing).
M.MAX = {
	[M.COFFEE] = 4,
	[M.PAPER]  = 6,
	[M.CAT]    = 2,
}

return M
