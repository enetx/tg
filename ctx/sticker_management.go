package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// CreateNewStickerSet represents a request to create a new sticker set.
type CreateNewStickerSet struct {
	ctx      *Context
	userID   int64
	name     String
	title    String
	stickers Slice[gotgbot.InputSticker]
	opts     *gotgbot.CreateNewStickerSetOpts
}

// StickerType sets the type of stickers in the set.
func (cns *CreateNewStickerSet) StickerType(stickerType String) *CreateNewStickerSet {
	cns.opts.StickerType = stickerType.Std()
	return cns
}

// NeedsRepainting marks stickers for repainting to custom emoji.
func (cns *CreateNewStickerSet) NeedsRepainting() *CreateNewStickerSet {
	cns.opts.NeedsRepainting = true
	return cns
}

// AddSticker adds a sticker to the new sticker set.
func (cns *CreateNewStickerSet) AddSticker(filename, format String, emojiList Slice[String]) *CreateNewStickerSet {
	sticker := gotgbot.InputSticker{
		Sticker:   filename.Std(),
		Format:    format.Std(),
		EmojiList: emojiList.ToStringSlice(),
	}

	cns.stickers.Push(sticker)

	return cns
}

// Keywords sets keywords for the last added sticker.
func (cns *CreateNewStickerSet) Keywords(keywords Slice[String]) *CreateNewStickerSet {
	if cns.stickers.NotEmpty() {
		cns.stickers[len(cns.stickers)-1].Keywords = keywords.ToStringSlice()
	}

	return cns
}

// MaskPosition sets the mask position for the last added sticker.
func (cns *CreateNewStickerSet) MaskPosition(point String, xShift, yShift, scale float64) *CreateNewStickerSet {
	if cns.stickers.NotEmpty() {
		cns.stickers[len(cns.stickers)-1].MaskPosition = &gotgbot.MaskPosition{
			Point:  point.Std(),
			XShift: xShift,
			YShift: yShift,
			Scale:  scale,
		}
	}

	return cns
}

// Timeout sets a custom timeout for this request.
func (cns *CreateNewStickerSet) Timeout(duration time.Duration) *CreateNewStickerSet {
	if cns.opts.RequestOpts == nil {
		cns.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cns.opts.RequestOpts.Timeout = duration

	return cns
}

// APIURL sets a custom API URL for this request.
func (cns *CreateNewStickerSet) APIURL(url String) *CreateNewStickerSet {
	if cns.opts.RequestOpts == nil {
		cns.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cns.opts.RequestOpts.APIURL = url.Std()

	return cns
}

// Send creates the new sticker set and returns the result.
func (cns *CreateNewStickerSet) Send() Result[bool] {
	if len(cns.stickers) == 0 {
		return Err[bool](Errorf("no stickers added to sticker set"))
	}

	return ResultOf(cns.ctx.Bot.Raw().
		CreateNewStickerSet(cns.userID, cns.name.Std(), cns.title.Std(), cns.stickers, cns.opts))
}

// AddStickerToSet represents a request to add a sticker to an existing set.
type AddStickerToSet struct {
	ctx     *Context
	userID  int64
	name    String
	sticker gotgbot.InputSticker
	opts    *gotgbot.AddStickerToSetOpts
}

// File sets the sticker file.
func (ats *AddStickerToSet) File(filename String) *AddStickerToSet {
	ats.sticker.Sticker = string(filename)
	return ats
}

// Format sets the sticker format.
func (ats *AddStickerToSet) Format(format String) *AddStickerToSet {
	ats.sticker.Format = format.Std()
	return ats
}

// EmojiList sets the emoji list for the sticker.
func (ats *AddStickerToSet) EmojiList(emojis Slice[String]) *AddStickerToSet {
	ats.sticker.EmojiList = emojis.ToStringSlice()
	return ats
}

// Keywords sets keywords for the sticker.
func (ats *AddStickerToSet) Keywords(keywords Slice[String]) *AddStickerToSet {
	ats.sticker.Keywords = keywords.ToStringSlice()
	return ats
}

// MaskPosition sets the mask position for the sticker.
func (ats *AddStickerToSet) MaskPosition(point String, xShift, yShift, scale float64) *AddStickerToSet {
	ats.sticker.MaskPosition = &gotgbot.MaskPosition{
		Point:  point.Std(),
		XShift: xShift,
		YShift: yShift,
		Scale:  scale,
	}

	return ats
}

// Timeout sets a custom timeout for this request.
func (ats *AddStickerToSet) Timeout(duration time.Duration) *AddStickerToSet {
	if ats.opts.RequestOpts == nil {
		ats.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ats.opts.RequestOpts.Timeout = duration

	return ats
}

// APIURL sets a custom API URL for this request.
func (ats *AddStickerToSet) APIURL(url String) *AddStickerToSet {
	if ats.opts.RequestOpts == nil {
		ats.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ats.opts.RequestOpts.APIURL = url.Std()

	return ats
}

// Send adds the sticker to the set and returns the result.
func (ats *AddStickerToSet) Send() Result[bool] {
	return ResultOf(ats.ctx.Bot.Raw().AddStickerToSet(ats.userID, ats.name.Std(), ats.sticker, ats.opts))
}

// GetStickerSet represents a request to get sticker set information.
type GetStickerSet struct {
	ctx  *Context
	name String
	opts *gotgbot.GetStickerSetOpts
}

// Timeout sets a custom timeout for this request.
func (gss *GetStickerSet) Timeout(duration time.Duration) *GetStickerSet {
	if gss.opts.RequestOpts == nil {
		gss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gss.opts.RequestOpts.Timeout = duration

	return gss
}

// APIURL sets a custom API URL for this request.
func (gss *GetStickerSet) APIURL(url String) *GetStickerSet {
	if gss.opts.RequestOpts == nil {
		gss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gss.opts.RequestOpts.APIURL = url.Std()

	return gss
}

// Send retrieves the sticker set information.
func (gss *GetStickerSet) Send() Result[*gotgbot.StickerSet] {
	return ResultOf(gss.ctx.Bot.Raw().GetStickerSet(gss.name.Std(), gss.opts))
}

// DeleteStickerSet represents a request to delete a sticker set.
type DeleteStickerSet struct {
	ctx  *Context
	name String
	opts *gotgbot.DeleteStickerSetOpts
}

// Timeout sets a custom timeout for this request.
func (dss *DeleteStickerSet) Timeout(duration time.Duration) *DeleteStickerSet {
	if dss.opts.RequestOpts == nil {
		dss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dss.opts.RequestOpts.Timeout = duration

	return dss
}

// APIURL sets a custom API URL for this request.
func (dss *DeleteStickerSet) APIURL(url String) *DeleteStickerSet {
	if dss.opts.RequestOpts == nil {
		dss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dss.opts.RequestOpts.APIURL = url.Std()

	return dss
}

// Send deletes the sticker set.
func (dss *DeleteStickerSet) Send() Result[bool] {
	return ResultOf(dss.ctx.Bot.Raw().DeleteStickerSet(dss.name.Std(), dss.opts))
}

// DeleteStickerFromSet represents a request to delete a sticker from a set.
type DeleteStickerFromSet struct {
	ctx     *Context
	sticker String
	opts    *gotgbot.DeleteStickerFromSetOpts
}

// Timeout sets a custom timeout for this request.
func (dsfs *DeleteStickerFromSet) Timeout(duration time.Duration) *DeleteStickerFromSet {
	if dsfs.opts.RequestOpts == nil {
		dsfs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dsfs.opts.RequestOpts.Timeout = duration

	return dsfs
}

// APIURL sets a custom API URL for this request.
func (dsfs *DeleteStickerFromSet) APIURL(url String) *DeleteStickerFromSet {
	if dsfs.opts.RequestOpts == nil {
		dsfs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dsfs.opts.RequestOpts.APIURL = url.Std()

	return dsfs
}

// Send deletes the sticker from the set.
func (dsfs *DeleteStickerFromSet) Send() Result[bool] {
	return ResultOf(dsfs.ctx.Bot.Raw().DeleteStickerFromSet(dsfs.sticker.Std(), dsfs.opts))
}

// SetStickerPositionInSet represents a request to set sticker position in set.
type SetStickerPositionInSet struct {
	ctx      *Context
	sticker  String
	position int64
	opts     *gotgbot.SetStickerPositionInSetOpts
}

// Timeout sets a custom timeout for this request.
func (sspis *SetStickerPositionInSet) Timeout(duration time.Duration) *SetStickerPositionInSet {
	if sspis.opts.RequestOpts == nil {
		sspis.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sspis.opts.RequestOpts.Timeout = duration

	return sspis
}

// APIURL sets a custom API URL for this request.
func (sspis *SetStickerPositionInSet) APIURL(url String) *SetStickerPositionInSet {
	if sspis.opts.RequestOpts == nil {
		sspis.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sspis.opts.RequestOpts.APIURL = url.Std()

	return sspis
}

// Send sets the sticker position in the set.
func (sspis *SetStickerPositionInSet) Send() Result[bool] {
	return ResultOf(sspis.ctx.Bot.Raw().SetStickerPositionInSet(sspis.sticker.Std(), int64(sspis.position), sspis.opts))
}

// SetStickerEmojiList represents a request to set sticker emoji list.
type SetStickerEmojiList struct {
	ctx       *Context
	sticker   String
	emojiList Slice[String]
	opts      *gotgbot.SetStickerEmojiListOpts
}

// EmojiList sets the emoji list for the sticker.
func (ssel *SetStickerEmojiList) EmojiList(emojis Slice[String]) *SetStickerEmojiList {
	ssel.emojiList = emojis
	return ssel
}

// Timeout sets a custom timeout for this request.
func (ssel *SetStickerEmojiList) Timeout(duration time.Duration) *SetStickerEmojiList {
	if ssel.opts.RequestOpts == nil {
		ssel.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssel.opts.RequestOpts.Timeout = duration

	return ssel
}

// APIURL sets a custom API URL for this request.
func (ssel *SetStickerEmojiList) APIURL(url String) *SetStickerEmojiList {
	if ssel.opts.RequestOpts == nil {
		ssel.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssel.opts.RequestOpts.APIURL = url.Std()

	return ssel
}

// Send sets the sticker emoji list.
func (ssel *SetStickerEmojiList) Send() Result[bool] {
	return ResultOf(ssel.ctx.Bot.Raw().
		SetStickerEmojiList(ssel.sticker.Std(), ssel.emojiList.ToStringSlice(), ssel.opts),
	)
}

// SetStickerKeywords represents a request to set sticker keywords.
type SetStickerKeywords struct {
	ctx      *Context
	sticker  String
	keywords Slice[String]
	opts     *gotgbot.SetStickerKeywordsOpts
}

// Keywords sets the keywords for the sticker.
func (ssk *SetStickerKeywords) Keywords(keywords Slice[String]) *SetStickerKeywords {
	ssk.keywords = keywords
	return ssk
}

// Timeout sets a custom timeout for this request.
func (ssk *SetStickerKeywords) Timeout(duration time.Duration) *SetStickerKeywords {
	if ssk.opts.RequestOpts == nil {
		ssk.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssk.opts.RequestOpts.Timeout = duration

	return ssk
}

// APIURL sets a custom API URL for this request.
func (ssk *SetStickerKeywords) APIURL(url String) *SetStickerKeywords {
	if ssk.opts.RequestOpts == nil {
		ssk.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssk.opts.RequestOpts.APIURL = url.Std()

	return ssk
}

// Send sets the sticker keywords.
func (ssk *SetStickerKeywords) Send() Result[bool] {
	ssk.opts.Keywords = ssk.keywords.ToStringSlice()
	return ResultOf(ssk.ctx.Bot.Raw().SetStickerKeywords(ssk.sticker.Std(), ssk.opts))
}

// SetStickerMaskPosition represents a request to set sticker mask position.
type SetStickerMaskPosition struct {
	ctx          *Context
	sticker      String
	maskPosition *gotgbot.MaskPosition
	opts         *gotgbot.SetStickerMaskPositionOpts
}

// MaskPosition sets the mask position for the sticker.
func (ssmp *SetStickerMaskPosition) MaskPosition(point String, xShift, yShift, scale float64) *SetStickerMaskPosition {
	ssmp.maskPosition = &gotgbot.MaskPosition{
		Point:  point.Std(),
		XShift: xShift,
		YShift: yShift,
		Scale:  scale,
	}

	return ssmp
}

// Timeout sets a custom timeout for this request.
func (ssmp *SetStickerMaskPosition) Timeout(duration time.Duration) *SetStickerMaskPosition {
	if ssmp.opts.RequestOpts == nil {
		ssmp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssmp.opts.RequestOpts.Timeout = duration

	return ssmp
}

// APIURL sets a custom API URL for this request.
func (ssmp *SetStickerMaskPosition) APIURL(url String) *SetStickerMaskPosition {
	if ssmp.opts.RequestOpts == nil {
		ssmp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssmp.opts.RequestOpts.APIURL = url.Std()

	return ssmp
}

// Send sets the sticker mask position.
func (ssmp *SetStickerMaskPosition) Send() Result[bool] {
	ssmp.opts.MaskPosition = ssmp.maskPosition
	return ResultOf(ssmp.ctx.Bot.Raw().SetStickerMaskPosition(ssmp.sticker.Std(), ssmp.opts))
}

// SetStickerSetThumbnail represents a request to set sticker set thumbnail.
type SetStickerSetThumbnail struct {
	ctx    *Context
	name   String
	userID int64
	format String
	opts   *gotgbot.SetStickerSetThumbnailOpts
	thumb  *File
	err    error
}

// Thumbnail sets the thumbnail file for the sticker set.
func (ssst *SetStickerSetThumbnail) Thumbnail(filename String) *SetStickerSetThumbnail {
	ssst.thumb = NewFile(filename)

	reader := ssst.thumb.Open()
	if reader.IsErr() {
		ssst.err = reader.Err()
		return ssst
	}

	ssst.opts.Thumbnail = gotgbot.InputFileByReader(ssst.thumb.Name().Std(), reader.Ok().Std())
	return ssst
}

// Format sets the thumbnail format.
// format of the thumbnail, must be one of "static" for a .WEBP or .PNG image,
// "animated" for a .TGS animation, or "video" for a .WEBM video.
func (ssst *SetStickerSetThumbnail) Format(format String) *SetStickerSetThumbnail {
	ssst.format = format
	return ssst
}

// Timeout sets a custom timeout for this request.
func (ssst *SetStickerSetThumbnail) Timeout(duration time.Duration) *SetStickerSetThumbnail {
	if ssst.opts.RequestOpts == nil {
		ssst.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssst.opts.RequestOpts.Timeout = duration

	return ssst
}

// APIURL sets a custom API URL for this request.
func (ssst *SetStickerSetThumbnail) APIURL(url String) *SetStickerSetThumbnail {
	if ssst.opts.RequestOpts == nil {
		ssst.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssst.opts.RequestOpts.APIURL = url.Std()

	return ssst
}

// Send sets the sticker set thumbnail.
func (ssst *SetStickerSetThumbnail) Send() Result[bool] {
	if ssst.err != nil {
		return Err[bool](ssst.err)
	}

	if ssst.thumb != nil {
		defer ssst.thumb.Close()
	}

	return ResultOf(ssst.ctx.Bot.Raw().
		SetStickerSetThumbnail(ssst.name.Std(), ssst.userID, ssst.format.Std(), ssst.opts),
	)
}

// UploadStickerFile represents a request to upload a sticker file.
type UploadStickerFile struct {
	ctx           *Context
	userID        int64
	sticker       gotgbot.InputFile
	stickerFormat String
	opts          *gotgbot.UploadStickerFileOpts
	file          *File
	err           error
}

// File sets the sticker file to upload.
func (usf *UploadStickerFile) File(filename String) *UploadStickerFile {
	usf.file = NewFile(filename)

	reader := usf.file.Open()
	if reader.IsErr() {
		usf.err = reader.Err()
		return usf
	}

	usf.sticker = gotgbot.InputFileByReader(usf.file.Name().Std(), reader.Ok().Std())
	return usf
}

// Format sets the sticker format.
func (usf *UploadStickerFile) Format(format String) *UploadStickerFile {
	usf.stickerFormat = format
	return usf
}

// Timeout sets a custom timeout for this request.
func (usf *UploadStickerFile) Timeout(duration time.Duration) *UploadStickerFile {
	if usf.opts.RequestOpts == nil {
		usf.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	usf.opts.RequestOpts.Timeout = duration

	return usf
}

// APIURL sets a custom API URL for this request.
func (usf *UploadStickerFile) APIURL(url String) *UploadStickerFile {
	if usf.opts.RequestOpts == nil {
		usf.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	usf.opts.RequestOpts.APIURL = url.Std()

	return usf
}

// Send uploads the sticker file.
func (usf *UploadStickerFile) Send() Result[*gotgbot.File] {
	if usf.err != nil {
		return Err[*gotgbot.File](usf.err)
	}

	if usf.file != nil {
		defer usf.file.Close()
	}

	return ResultOf(usf.ctx.Bot.Raw().UploadStickerFile(usf.userID, usf.sticker, usf.stickerFormat.Std(), usf.opts))
}

// GetCustomEmojiStickers represents a request to get custom emoji stickers.
type GetCustomEmojiStickers struct {
	ctx            *Context
	customEmojiIDs Slice[String]
	opts           *gotgbot.GetCustomEmojiStickersOpts
}

// Timeout sets a custom timeout for this request.
func (gces *GetCustomEmojiStickers) Timeout(duration time.Duration) *GetCustomEmojiStickers {
	if gces.opts.RequestOpts == nil {
		gces.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gces.opts.RequestOpts.Timeout = duration

	return gces
}

// APIURL sets a custom API URL for this request.
func (gces *GetCustomEmojiStickers) APIURL(url String) *GetCustomEmojiStickers {
	if gces.opts.RequestOpts == nil {
		gces.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gces.opts.RequestOpts.APIURL = url.Std()

	return gces
}

// Send retrieves the custom emoji stickers.
func (gces *GetCustomEmojiStickers) Send() Result[Slice[gotgbot.Sticker]] {
	stickers, err := gces.ctx.Bot.Raw().GetCustomEmojiStickers(gces.customEmojiIDs.ToStringSlice(), gces.opts)
	return ResultOf[Slice[gotgbot.Sticker]](stickers, err)
}
