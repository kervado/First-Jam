-- Persistent game settings: load / save on the local machine and apply them
-- to the running engine. Shared by main.script (on startup) and the settings
-- GUI (on change).

local M = {}

M.DEFAULTS = { fullscreen = true, volume = 0.5 }

local function save_path()
	return sys.get_save_file("JamOfficePsychosis", "settings")
end

local function clamp01(v)
	if v < 0 then return 0 end
	if v > 1 then return 1 end
	return v
end

-- Returns a settings table, filled with defaults for anything missing/invalid.
function M.load()
	local t = { fullscreen = M.DEFAULTS.fullscreen, volume = M.DEFAULTS.volume }
	local data = sys.load(save_path()) -- {} when the file does not exist yet
	if data then
		if type(data.fullscreen) == "boolean" then
			t.fullscreen = data.fullscreen
		end
		if type(data.volume) == "number" then
			t.volume = clamp01(data.volume)
		end
	end
	return t
end

-- Writes the settings to disk. Returns true on success.
function M.save(t)
	return sys.save(save_path(), {
		fullscreen = t.fullscreen and true or false,
		volume = clamp01(t.volume or M.DEFAULTS.volume),
	})
end

-- Applies the settings to the engine (window mode + master audio gain).
function M.apply(t)
	if sound and sound.set_group_gain then
		sound.set_group_gain(hash("master"), clamp01(t.volume or M.DEFAULTS.volume))
	end
	-- Runtime fullscreen toggling needs the DefOS native extension; the
	-- built-in window module has no fullscreen API.
	if defos and defos.set_fullscreen then
		defos.set_fullscreen(t.fullscreen and true or false)
	end
end

return M
