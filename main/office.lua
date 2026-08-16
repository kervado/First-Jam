--------------------------------------------------------------------------------
-- office.lua
--
-- Level-wide shared state and helpers. With shared_state = 1 this module is a
-- single shared table, so `office.psychosis` is the one common variable for the
-- whole level (the controller writes it, solid objects read it).
--------------------------------------------------------------------------------

local M = {}

M.ROOM_W = 1920
M.ROOM_H = 1080

-- the level psychosis meter, 0..100 (reset by the controller on level start)
M.psychosis = 0

-- The floor always sits just behind every other object (a hair inside the
-- near plane so it is never clipped).
M.Z_FLOOR = -0.98

-- Fake-isometric depth: an object's Z is derived from its Y, in the range
-- [-1 .. 1]. Objects lower on the screen (small Y) are closest to the viewer
-- and must be drawn in front; the floor stays behind everything.
--
--   y = 0       (bottom of the play area) -> Z = +1.0  (drawn in front)
--   y = ROOM_H  (top of the play area)    -> Z = -1.0  (drawn behind)
--
-- IMPORTANT: Defold's default renderer draws a HIGHER Z on top, which is the
-- opposite of the near/far-plane intuition (where far_z = +1 would be "back").
-- So to put the floor at the back and near-camera objects in front, the sign is
-- the reverse of "small Y -> small Z". If you truly need the literal values
-- from the spec (bottom = -1, floor = +1), swap to `return t * 2 - 1` and set
-- Z_FLOOR = 1.0 -- but then higher-Z-on-top will render the layering inverted.
function M.z_for_y(y)
	local t = y / M.ROOM_H
	if t < 0 then t = 0 elseif t > 1 then t = 1 end
	local z = 1 - t * 2
	-- keep headroom inside [-1, 1] so per-child / per-component offsets (e.g. a
	-- request icon sitting at local z +0.1 on a desk) never cross the far plane,
	-- and so objects stay in front of the floor at Z_FLOOR
	if z > 0.85 then z = 0.85 elseif z < -0.85 then z = -0.85 end
	return z
end

-- Set the calling object's Z from its WORLD Y (fake-isometric depth). Local X/Y
-- are kept, and any parent's Z is compensated so the resulting WORLD Z is what
-- z_for_y expects (objects may be children of other objects). Pass `world_z` to
-- force a specific world Z (used by the floor). Returns the world Z applied.
function M.apply_depth(world_z)
	local wp = go.get_world_position()
	local lp = go.get_position()
	local target = world_z or M.z_for_y(wp.y)
	local parent_z = wp.z - lp.z           -- the parent's world-Z contribution
	go.set_position(vmath.vector3(lp.x, lp.y, target - parent_z))
	return target
end

return M
