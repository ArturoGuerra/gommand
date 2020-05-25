# Handling deleted messages

For some bots, being able to track deleted messages in an easy to use way is important due to the ability to get the content/author of deleted messages. In standalone disgord, you need to manually cache messages. However, gommand has the ability built in to cache deleted messages. To use this handler, simply set the `DeletedMessageHandler` attribute of the router configuration to a struct of the type `&gommand.DeletedMessageHandler`. You can then set the following attributes in this handler:

- `Limit`: Defines the maximum amount of cached messages. -1 = unlimited (not suggested if it's in-memory since it'll lead to memory leaks), 0 = default, >0 = user set maximum. This should run on a First In First Out (FIFO) basis. By default, this will be set to 1,000 messages. Messages which have been purged due to this limit will not have an event fired for them.
- `Callback`: The callback of type `func(s disgord.Session, msg *disgord.Message)` which is called when a message is deleted. As with commands, for ease of use the `Member` attribute is set on the message.
- `MessageCacheStorageAdapter`: The [message cache storage adapter](/message-cache-storage-adapter) which is used for this. If this is not set, it will default to the built-in in-memory caching adapter.

## Message cache storage adapter
By default (like other libraries such as discord.py), gommand keeps a finite amount of messages cached into RAM which is set by the user in the deleted message handler parameters (or defaults to 1,000). However, if you wish to store these in a database somewhere (normally for memory management purposes), you will likely want to want to write your own message caching adapter. In gommand, memory cachers use the `gommand.MemoryCacheStorageAdapter` interface. This contains the following functions which need to be set:

- `Init()`: Called on the initialisation of the router.
- `GetAndDelete(ChannelID, MessageID disgord.Snowflake) *disgord.Message`: Gets a message from the cache and then deletes it since this is only called when the message is being deleted so it will then be unneeded.
- `Delete(ChannelID, MessageID disgord.Snowflake)`: Deletes a message from the cache.
- `DeleteChannelsMessages(ChannelID disgord.Snowflake)`: Deletes all messages cached for a specific channel.
- `Set(ChannelID, MessageID disgord.Snowflake, Message *disgord.Message, Limit uint)`: Sets an item in the cache. The limit is passed through so that you can implement a simple First In First Out (FIFO) caching system. The limit will be 0 if it is set to unlimited.

The following manage storing channel/guild ID relationships. This is important so that if a guild is removed, we know what channel ID's to purge from the cache:

- `GetAllChannelIDs(GuildID disgord.Snowflake) []disgord.Snowflake`: Get all channel ID's which have a relationship with a specific guild ID.
- `AddChannelID(GuildID, ChannelID disgord.Snowflake)`: Add a relationship between a guild ID and a channel ID.
- `RemoveChannelID(GuildID, ChannelID disgord.Snowflake)`: Remove a channel ID's relationship with a guild ID.
- `RemoveGuild(GuildID disgord.Snowflake)`: Remove all channel ID relationships with a specific guild ID.