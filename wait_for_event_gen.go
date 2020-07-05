// Code generated by gommand. DO NOT EDIT.
package gommand

import (
	"context"
	"github.com/andersfylling/disgord"
	"time"
)

// WaitManager is used to manage waiting within the context.
type WaitManager struct {
    ctx *Context
}

// This allows you to wait for a event with specific conditions. You should NOT block during the check function.
func (w *WaitManager) waitForEvent(ctx context.Context, EventName string, CheckFunc func(s disgord.Session, evt interface{}) bool) interface{} {
	x := make(chan interface{})
	middleware := func(evt interface{}) interface{} {
		if !CheckFunc(w.ctx.Session, evt) {
			return nil
		}
		return evt
	}
	var timer *time.Timer
	until, ok := ctx.Deadline()
	if ok {
		timer = time.AfterFunc(time.Until(until), func() {
			x <- nil
		})
	}
	handleEmit := func(e interface{}) {
		x <- e
		if timer != nil {
			timer.Stop()
		}
	}
	var emitChan interface{}
	switch EventName {
	case disgord.EvtChannelCreate:
		emitChan = func(_ disgord.Session, e *disgord.ChannelCreate) { handleEmit(e) }
	case disgord.EvtChannelUpdate:
		emitChan = func(_ disgord.Session, e *disgord.ChannelUpdate) { handleEmit(e) }
	case disgord.EvtChannelDelete:
		emitChan = func(_ disgord.Session, e *disgord.ChannelDelete) { handleEmit(e) }
	case disgord.EvtChannelPinsUpdate:
		emitChan = func(_ disgord.Session, e *disgord.ChannelPinsUpdate) { handleEmit(e) }
	case disgord.EvtTypingStart:
		emitChan = func(_ disgord.Session, e *disgord.TypingStart) { handleEmit(e) }
	case disgord.EvtInviteDelete:
		emitChan = func(_ disgord.Session, e *disgord.InviteDelete) { handleEmit(e) }
	case disgord.EvtMessageCreate:
		emitChan = func(_ disgord.Session, e *disgord.MessageCreate) { handleEmit(e) }
	case disgord.EvtMessageUpdate:
		emitChan = func(_ disgord.Session, e *disgord.MessageUpdate) { handleEmit(e) }
	case disgord.EvtMessageDelete:
		emitChan = func(_ disgord.Session, e *disgord.MessageDelete) { handleEmit(e) }
	case disgord.EvtMessageDeleteBulk:
		emitChan = func(_ disgord.Session, e *disgord.MessageDeleteBulk) { handleEmit(e) }
	case disgord.EvtMessageReactionAdd:
		emitChan = func(_ disgord.Session, e *disgord.MessageReactionAdd) { handleEmit(e) }
	case disgord.EvtMessageReactionRemove:
		emitChan = func(_ disgord.Session, e *disgord.MessageReactionRemove) { handleEmit(e) }
	case disgord.EvtMessageReactionRemoveAll:
		emitChan = func(_ disgord.Session, e *disgord.MessageReactionRemoveAll) { handleEmit(e) }
	case disgord.EvtGuildEmojisUpdate:
		emitChan = func(_ disgord.Session, e *disgord.GuildEmojisUpdate) { handleEmit(e) }
	case disgord.EvtGuildCreate:
		emitChan = func(_ disgord.Session, e *disgord.GuildCreate) { handleEmit(e) }
	case disgord.EvtGuildUpdate:
		emitChan = func(_ disgord.Session, e *disgord.GuildUpdate) { handleEmit(e) }
	case disgord.EvtGuildDelete:
		emitChan = func(_ disgord.Session, e *disgord.GuildDelete) { handleEmit(e) }
	case disgord.EvtGuildBanAdd:
		emitChan = func(_ disgord.Session, e *disgord.GuildBanAdd) { handleEmit(e) }
	case disgord.EvtGuildBanRemove:
		emitChan = func(_ disgord.Session, e *disgord.GuildBanRemove) { handleEmit(e) }
	case disgord.EvtGuildMemberAdd:
		emitChan = func(_ disgord.Session, e *disgord.GuildMemberAdd) { handleEmit(e) }
	case disgord.EvtGuildMemberRemove:
		emitChan = func(_ disgord.Session, e *disgord.GuildMemberRemove) { handleEmit(e) }
	case disgord.EvtGuildMemberUpdate:
		emitChan = func(_ disgord.Session, e *disgord.GuildMemberUpdate) { handleEmit(e) }
	case disgord.EvtGuildRoleCreate:
		emitChan = func(_ disgord.Session, e *disgord.GuildRoleCreate) { handleEmit(e) }
	case disgord.EvtGuildRoleUpdate:
		emitChan = func(_ disgord.Session, e *disgord.GuildRoleUpdate) { handleEmit(e) }
	case disgord.EvtGuildRoleDelete:
		emitChan = func(_ disgord.Session, e *disgord.GuildRoleDelete) { handleEmit(e) }
	case disgord.EvtPresenceUpdate:
		emitChan = func(_ disgord.Session, e *disgord.PresenceUpdate) { handleEmit(e) }
	case disgord.EvtUserUpdate:
		emitChan = func(_ disgord.Session, e *disgord.UserUpdate) { handleEmit(e) }
	case disgord.EvtVoiceStateUpdate:
		emitChan = func(_ disgord.Session, e *disgord.VoiceStateUpdate) { handleEmit(e) }
	case disgord.EvtVoiceServerUpdate:
		emitChan = func(_ disgord.Session, e *disgord.VoiceServerUpdate) { handleEmit(e) }
	case disgord.EvtWebhooksUpdate:
		emitChan = func(_ disgord.Session, e *disgord.WebhooksUpdate) { handleEmit(e) }
	case disgord.EvtInviteCreate:
		emitChan = func(_ disgord.Session, e *disgord.InviteCreate) { handleEmit(e) }
	default:
		panic("unknown event")
	}
	w.ctx.Session.On(EventName, middleware, emitChan, &disgord.Ctrl{Runs: 1, Until: until})
	return <-x
}

// WaitForChannelCreate allows you to wait for the ChannelCreate event. You should NOT block during the check function.
func (w *WaitManager) WaitForChannelCreate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.ChannelCreate) bool) *disgord.ChannelCreate {
	x := w.waitForEvent(ctx, disgord.EvtChannelCreate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.ChannelCreate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.ChannelCreate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForChannelUpdate allows you to wait for the ChannelUpdate event. You should NOT block during the check function.
func (w *WaitManager) WaitForChannelUpdate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.ChannelUpdate) bool) *disgord.ChannelUpdate {
	x := w.waitForEvent(ctx, disgord.EvtChannelUpdate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.ChannelUpdate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.ChannelUpdate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForChannelDelete allows you to wait for the ChannelDelete event. You should NOT block during the check function.
func (w *WaitManager) WaitForChannelDelete(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.ChannelDelete) bool) *disgord.ChannelDelete {
	x := w.waitForEvent(ctx, disgord.EvtChannelDelete, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.ChannelDelete); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.ChannelDelete)
	if !ok {
	    return nil
	}
	return r
}

// WaitForChannelPinsUpdate allows you to wait for the ChannelPinsUpdate event. You should NOT block during the check function.
func (w *WaitManager) WaitForChannelPinsUpdate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.ChannelPinsUpdate) bool) *disgord.ChannelPinsUpdate {
	x := w.waitForEvent(ctx, disgord.EvtChannelPinsUpdate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.ChannelPinsUpdate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.ChannelPinsUpdate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForTypingStart allows you to wait for the TypingStart event. You should NOT block during the check function.
func (w *WaitManager) WaitForTypingStart(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.TypingStart) bool) *disgord.TypingStart {
	x := w.waitForEvent(ctx, disgord.EvtTypingStart, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.TypingStart); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.TypingStart)
	if !ok {
	    return nil
	}
	return r
}

// WaitForInviteDelete allows you to wait for the InviteDelete event. You should NOT block during the check function.
func (w *WaitManager) WaitForInviteDelete(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.InviteDelete) bool) *disgord.InviteDelete {
	x := w.waitForEvent(ctx, disgord.EvtInviteDelete, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.InviteDelete); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.InviteDelete)
	if !ok {
	    return nil
	}
	return r
}

// WaitForMessageCreate allows you to wait for the MessageCreate event. You should NOT block during the check function.
func (w *WaitManager) WaitForMessageCreate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.MessageCreate) bool) *disgord.MessageCreate {
	x := w.waitForEvent(ctx, disgord.EvtMessageCreate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.MessageCreate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.MessageCreate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForMessageUpdate allows you to wait for the MessageUpdate event. You should NOT block during the check function.
func (w *WaitManager) WaitForMessageUpdate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.MessageUpdate) bool) *disgord.MessageUpdate {
	x := w.waitForEvent(ctx, disgord.EvtMessageUpdate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.MessageUpdate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.MessageUpdate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForMessageDelete allows you to wait for the MessageDelete event. You should NOT block during the check function.
func (w *WaitManager) WaitForMessageDelete(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.MessageDelete) bool) *disgord.MessageDelete {
	x := w.waitForEvent(ctx, disgord.EvtMessageDelete, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.MessageDelete); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.MessageDelete)
	if !ok {
	    return nil
	}
	return r
}

// WaitForMessageDeleteBulk allows you to wait for the MessageDeleteBulk event. You should NOT block during the check function.
func (w *WaitManager) WaitForMessageDeleteBulk(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.MessageDeleteBulk) bool) *disgord.MessageDeleteBulk {
	x := w.waitForEvent(ctx, disgord.EvtMessageDeleteBulk, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.MessageDeleteBulk); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.MessageDeleteBulk)
	if !ok {
	    return nil
	}
	return r
}

// WaitForMessageReactionAdd allows you to wait for the MessageReactionAdd event. You should NOT block during the check function.
func (w *WaitManager) WaitForMessageReactionAdd(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.MessageReactionAdd) bool) *disgord.MessageReactionAdd {
	x := w.waitForEvent(ctx, disgord.EvtMessageReactionAdd, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.MessageReactionAdd); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.MessageReactionAdd)
	if !ok {
	    return nil
	}
	return r
}

// WaitForMessageReactionRemove allows you to wait for the MessageReactionRemove event. You should NOT block during the check function.
func (w *WaitManager) WaitForMessageReactionRemove(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.MessageReactionRemove) bool) *disgord.MessageReactionRemove {
	x := w.waitForEvent(ctx, disgord.EvtMessageReactionRemove, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.MessageReactionRemove); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.MessageReactionRemove)
	if !ok {
	    return nil
	}
	return r
}

// WaitForMessageReactionRemoveAll allows you to wait for the MessageReactionRemoveAll event. You should NOT block during the check function.
func (w *WaitManager) WaitForMessageReactionRemoveAll(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.MessageReactionRemoveAll) bool) *disgord.MessageReactionRemoveAll {
	x := w.waitForEvent(ctx, disgord.EvtMessageReactionRemoveAll, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.MessageReactionRemoveAll); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.MessageReactionRemoveAll)
	if !ok {
	    return nil
	}
	return r
}

// WaitForGuildEmojisUpdate allows you to wait for the GuildEmojisUpdate event. You should NOT block during the check function.
func (w *WaitManager) WaitForGuildEmojisUpdate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.GuildEmojisUpdate) bool) *disgord.GuildEmojisUpdate {
	x := w.waitForEvent(ctx, disgord.EvtGuildEmojisUpdate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.GuildEmojisUpdate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.GuildEmojisUpdate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForGuildCreate allows you to wait for the GuildCreate event. You should NOT block during the check function.
func (w *WaitManager) WaitForGuildCreate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.GuildCreate) bool) *disgord.GuildCreate {
	x := w.waitForEvent(ctx, disgord.EvtGuildCreate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.GuildCreate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.GuildCreate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForGuildUpdate allows you to wait for the GuildUpdate event. You should NOT block during the check function.
func (w *WaitManager) WaitForGuildUpdate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.GuildUpdate) bool) *disgord.GuildUpdate {
	x := w.waitForEvent(ctx, disgord.EvtGuildUpdate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.GuildUpdate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.GuildUpdate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForGuildDelete allows you to wait for the GuildDelete event. You should NOT block during the check function.
func (w *WaitManager) WaitForGuildDelete(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.GuildDelete) bool) *disgord.GuildDelete {
	x := w.waitForEvent(ctx, disgord.EvtGuildDelete, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.GuildDelete); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.GuildDelete)
	if !ok {
	    return nil
	}
	return r
}

// WaitForGuildBanAdd allows you to wait for the GuildBanAdd event. You should NOT block during the check function.
func (w *WaitManager) WaitForGuildBanAdd(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.GuildBanAdd) bool) *disgord.GuildBanAdd {
	x := w.waitForEvent(ctx, disgord.EvtGuildBanAdd, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.GuildBanAdd); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.GuildBanAdd)
	if !ok {
	    return nil
	}
	return r
}

// WaitForGuildBanRemove allows you to wait for the GuildBanRemove event. You should NOT block during the check function.
func (w *WaitManager) WaitForGuildBanRemove(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.GuildBanRemove) bool) *disgord.GuildBanRemove {
	x := w.waitForEvent(ctx, disgord.EvtGuildBanRemove, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.GuildBanRemove); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.GuildBanRemove)
	if !ok {
	    return nil
	}
	return r
}

// WaitForGuildMemberAdd allows you to wait for the GuildMemberAdd event. You should NOT block during the check function.
func (w *WaitManager) WaitForGuildMemberAdd(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.GuildMemberAdd) bool) *disgord.GuildMemberAdd {
	x := w.waitForEvent(ctx, disgord.EvtGuildMemberAdd, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.GuildMemberAdd); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.GuildMemberAdd)
	if !ok {
	    return nil
	}
	return r
}

// WaitForGuildMemberRemove allows you to wait for the GuildMemberRemove event. You should NOT block during the check function.
func (w *WaitManager) WaitForGuildMemberRemove(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.GuildMemberRemove) bool) *disgord.GuildMemberRemove {
	x := w.waitForEvent(ctx, disgord.EvtGuildMemberRemove, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.GuildMemberRemove); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.GuildMemberRemove)
	if !ok {
	    return nil
	}
	return r
}

// WaitForGuildMemberUpdate allows you to wait for the GuildMemberUpdate event. You should NOT block during the check function.
func (w *WaitManager) WaitForGuildMemberUpdate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.GuildMemberUpdate) bool) *disgord.GuildMemberUpdate {
	x := w.waitForEvent(ctx, disgord.EvtGuildMemberUpdate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.GuildMemberUpdate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.GuildMemberUpdate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForGuildRoleCreate allows you to wait for the GuildRoleCreate event. You should NOT block during the check function.
func (w *WaitManager) WaitForGuildRoleCreate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.GuildRoleCreate) bool) *disgord.GuildRoleCreate {
	x := w.waitForEvent(ctx, disgord.EvtGuildRoleCreate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.GuildRoleCreate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.GuildRoleCreate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForGuildRoleUpdate allows you to wait for the GuildRoleUpdate event. You should NOT block during the check function.
func (w *WaitManager) WaitForGuildRoleUpdate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.GuildRoleUpdate) bool) *disgord.GuildRoleUpdate {
	x := w.waitForEvent(ctx, disgord.EvtGuildRoleUpdate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.GuildRoleUpdate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.GuildRoleUpdate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForGuildRoleDelete allows you to wait for the GuildRoleDelete event. You should NOT block during the check function.
func (w *WaitManager) WaitForGuildRoleDelete(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.GuildRoleDelete) bool) *disgord.GuildRoleDelete {
	x := w.waitForEvent(ctx, disgord.EvtGuildRoleDelete, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.GuildRoleDelete); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.GuildRoleDelete)
	if !ok {
	    return nil
	}
	return r
}

// WaitForPresenceUpdate allows you to wait for the PresenceUpdate event. You should NOT block during the check function.
func (w *WaitManager) WaitForPresenceUpdate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.PresenceUpdate) bool) *disgord.PresenceUpdate {
	x := w.waitForEvent(ctx, disgord.EvtPresenceUpdate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.PresenceUpdate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.PresenceUpdate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForUserUpdate allows you to wait for the UserUpdate event. You should NOT block during the check function.
func (w *WaitManager) WaitForUserUpdate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.UserUpdate) bool) *disgord.UserUpdate {
	x := w.waitForEvent(ctx, disgord.EvtUserUpdate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.UserUpdate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.UserUpdate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForVoiceStateUpdate allows you to wait for the VoiceStateUpdate event. You should NOT block during the check function.
func (w *WaitManager) WaitForVoiceStateUpdate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.VoiceStateUpdate) bool) *disgord.VoiceStateUpdate {
	x := w.waitForEvent(ctx, disgord.EvtVoiceStateUpdate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.VoiceStateUpdate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.VoiceStateUpdate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForVoiceServerUpdate allows you to wait for the VoiceServerUpdate event. You should NOT block during the check function.
func (w *WaitManager) WaitForVoiceServerUpdate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.VoiceServerUpdate) bool) *disgord.VoiceServerUpdate {
	x := w.waitForEvent(ctx, disgord.EvtVoiceServerUpdate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.VoiceServerUpdate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.VoiceServerUpdate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForWebhooksUpdate allows you to wait for the WebhooksUpdate event. You should NOT block during the check function.
func (w *WaitManager) WaitForWebhooksUpdate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.WebhooksUpdate) bool) *disgord.WebhooksUpdate {
	x := w.waitForEvent(ctx, disgord.EvtWebhooksUpdate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.WebhooksUpdate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.WebhooksUpdate)
	if !ok {
	    return nil
	}
	return r
}

// WaitForInviteCreate allows you to wait for the InviteCreate event. You should NOT block during the check function.
func (w *WaitManager) WaitForInviteCreate(ctx context.Context, CheckFunc func(s disgord.Session, evt *disgord.InviteCreate) bool) *disgord.InviteCreate {
	x := w.waitForEvent(ctx, disgord.EvtInviteCreate, func(s disgord.Session, evt interface{}) bool {
		if e, ok := evt.(*disgord.InviteCreate); ok {
			return CheckFunc(s, e)
		}
		return false
	})
	r, ok := x.(*disgord.InviteCreate)
	if !ok {
	    return nil
	}
	return r
}