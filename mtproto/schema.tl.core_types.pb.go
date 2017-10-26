// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.tl.core_types.proto

/*
Package mtproto is a generated protocol buffer package.

It is generated from these files:
	schema.tl.core_types.proto
	schema.tl.crc32.proto
	schema.tl.handshake.proto
	schema.tl.sync.proto
	schema.tl.sync.service.proto
	schema.tl.transport.proto

It has these top-level messages:
	Bool
	True
	Error
	Null
	TLBoolFalse
	TLBoolTrue
	TLTrue
	TLError
	TLNull
	ResPQ
	P_QInnerData
	Server_DH_Params
	Server_DHInnerData
	Client_DH_Inner_Data
	SetClient_DHParamsAnswer
	DestroyAuthKeyRes
	TLResPQ
	TLPQInnerData
	TLServer_DHParamsFail
	TLServer_DHParamsOk
	TLServer_DHInnerData
	TLClient_DHInnerData
	TLDhGenOk
	TLDhGenRetry
	TLDhGenFail
	TLDestroyAuthKeyOk
	TLDestroyAuthKeyNone
	TLDestroyAuthKeyFail
	TLReqPq
	TLReq_DHParams
	TLSetClient_DHParams
	TLDestroyAuthKey
	InputPeer
	InputUser
	InputContact
	InputFile
	InputMedia
	InputChatPhoto
	InputGeoPoint
	InputPhoto
	InputFileLocation
	InputAppEvent
	Peer
	Storage_FileType
	FileLocation
	User
	UserProfilePhoto
	UserStatus
	Chat
	ChatFull
	ChatParticipant
	ChatParticipants
	ChatPhoto
	Message
	MessageMedia
	MessageAction
	Dialog
	Photo
	PhotoSize
	GeoPoint
	Auth_CheckedPhone
	Auth_SentCode
	Auth_Authorization
	Auth_ExportedAuthorization
	InputNotifyPeer
	InputPeerNotifyEvents
	InputPeerNotifySettings
	PeerNotifyEvents
	PeerNotifySettings
	PeerSettings
	WallPaper
	ReportReason
	UserFull
	Contact
	ImportedContact
	ContactBlocked
	ContactStatus
	Contacts_Link
	Contacts_Contacts
	Contacts_ImportedContacts
	Contacts_Blocked
	Messages_Dialogs
	Messages_Messages
	Messages_Chats
	Messages_ChatFull
	Messages_AffectedHistory
	MessagesFilter
	Update
	Updates_State
	Updates_Difference
	Updates
	Photos_Photos
	Photos_Photo
	Upload_File
	DcOption
	Config
	NearestDc
	Help_AppUpdate
	Help_InviteText
	EncryptedChat
	InputEncryptedChat
	EncryptedFile
	InputEncryptedFile
	EncryptedMessage
	Messages_DhConfig
	Messages_SentEncryptedMessage
	InputDocument
	Document
	Help_Support
	NotifyPeer
	SendMessageAction
	Contacts_Found
	InputPrivacyKey
	PrivacyKey
	InputPrivacyRule
	PrivacyRule
	Account_PrivacyRules
	AccountDaysTTL
	DocumentAttribute
	Messages_Stickers
	StickerPack
	Messages_AllStickers
	DisabledFeature
	Messages_AffectedMessages
	ContactLink
	WebPage
	Authorization
	Account_Authorizations
	Account_Password
	Account_PasswordSettings
	Account_PasswordInputSettings
	Auth_PasswordRecovery
	ReceivedNotifyMessage
	ExportedChatInvite
	ChatInvite
	InputStickerSet
	StickerSet
	Messages_StickerSet
	BotCommand
	BotInfo
	KeyboardButton
	KeyboardButtonRow
	ReplyMarkup
	MessageEntity
	InputChannel
	Contacts_ResolvedPeer
	MessageRange
	Updates_ChannelDifference
	ChannelMessagesFilter
	ChannelParticipant
	ChannelParticipantsFilter
	Channels_ChannelParticipants
	Channels_ChannelParticipant
	Help_TermsOfService
	FoundGif
	Messages_FoundGifs
	Messages_SavedGifs
	InputBotInlineMessage
	InputBotInlineResult
	BotInlineMessage
	BotInlineResult
	Messages_BotResults
	ExportedMessageLink
	MessageFwdHeader
	Auth_CodeType
	Auth_SentCodeType
	Messages_BotCallbackAnswer
	Messages_MessageEditData
	InputBotInlineMessageID
	InlineBotSwitchPM
	Messages_PeerDialogs
	TopPeer
	TopPeerCategory
	TopPeerCategoryPeers
	Contacts_TopPeers
	DraftMessage
	Messages_FeaturedStickers
	Messages_RecentStickers
	Messages_ArchivedStickers
	Messages_StickerSetInstallResult
	StickerSetCovered
	MaskCoords
	InputStickeredMedia
	Game
	InputGame
	HighScore
	Messages_HighScores
	RichText
	PageBlock
	Page
	PhoneCallDiscardReason
	DataJSON
	LabeledPrice
	Invoice
	PaymentCharge
	PostAddress
	PaymentRequestedInfo
	PaymentSavedCredentials
	WebDocument
	InputWebDocument
	InputWebFileLocation
	Upload_WebFile
	Payments_PaymentForm
	Payments_ValidatedRequestedInfo
	Payments_PaymentResult
	Payments_PaymentReceipt
	Payments_SavedInfo
	InputPaymentCredentials
	Account_TmpPassword
	ShippingOption
	InputStickerSetItem
	InputPhoneCall
	PhoneCall
	PhoneConnection
	PhoneCallProtocol
	Phone_PhoneCall
	Upload_CdnFile
	CdnPublicKey
	CdnConfig
	LangPackString
	LangPackDifference
	LangPackLanguage
	ChannelAdminRights
	ChannelBannedRights
	ChannelAdminLogEventAction
	ChannelAdminLogEvent
	Channels_AdminLogResults
	ChannelAdminLogEventsFilter
	PopularContact
	CdnFileHash
	Messages_FavedStickers
	TLInputPeerEmpty
	TLInputPeerSelf
	TLInputPeerChat
	TLInputPeerUser
	TLInputPeerChannel
	TLInputUserEmpty
	TLInputUserSelf
	TLInputUser
	TLInputPhoneContact
	TLInputFile
	TLInputFileBig
	TLInputMediaEmpty
	TLInputMediaUploadedPhoto
	TLInputMediaPhoto
	TLInputMediaGeoPoint
	TLInputMediaContact
	TLInputMediaUploadedDocument
	TLInputMediaDocument
	TLInputMediaVenue
	TLInputMediaGifExternal
	TLInputMediaPhotoExternal
	TLInputMediaDocumentExternal
	TLInputMediaGame
	TLInputMediaInvoice
	TLInputChatPhotoEmpty
	TLInputChatUploadedPhoto
	TLInputChatPhoto
	TLInputGeoPointEmpty
	TLInputGeoPoint
	TLInputPhotoEmpty
	TLInputPhoto
	TLInputFileLocation
	TLInputEncryptedFileLocation
	TLInputDocumentFileLocation
	TLInputAppEvent
	TLPeerUser
	TLPeerChat
	TLPeerChannel
	TLStorageFileUnknown
	TLStorageFilePartial
	TLStorageFileJpeg
	TLStorageFileGif
	TLStorageFilePng
	TLStorageFilePdf
	TLStorageFileMp3
	TLStorageFileMov
	TLStorageFileMp4
	TLStorageFileWebp
	TLFileLocationUnavailable
	TLFileLocation
	TLUserEmpty
	TLUser
	TLUserProfilePhotoEmpty
	TLUserProfilePhoto
	TLUserStatusEmpty
	TLUserStatusOnline
	TLUserStatusOffline
	TLUserStatusRecently
	TLUserStatusLastWeek
	TLUserStatusLastMonth
	TLChatEmpty
	TLChat
	TLChatForbidden
	TLChannel
	TLChannelForbidden
	TLChatFull
	TLChannelFull
	TLChatParticipant
	TLChatParticipantCreator
	TLChatParticipantAdmin
	TLChatParticipantsForbidden
	TLChatParticipants
	TLChatPhotoEmpty
	TLChatPhoto
	TLMessageEmpty
	TLMessage
	TLMessageService
	TLMessageMediaEmpty
	TLMessageMediaPhoto
	TLMessageMediaGeo
	TLMessageMediaContact
	TLMessageMediaUnsupported
	TLMessageMediaDocument
	TLMessageMediaWebPage
	TLMessageMediaVenue
	TLMessageMediaGame
	TLMessageMediaInvoice
	TLMessageActionEmpty
	TLMessageActionChatCreate
	TLMessageActionChatEditTitle
	TLMessageActionChatEditPhoto
	TLMessageActionChatDeletePhoto
	TLMessageActionChatAddUser
	TLMessageActionChatDeleteUser
	TLMessageActionChatJoinedByLink
	TLMessageActionChannelCreate
	TLMessageActionChatMigrateTo
	TLMessageActionChannelMigrateFrom
	TLMessageActionPinMessage
	TLMessageActionHistoryClear
	TLMessageActionGameScore
	TLMessageActionPaymentSentMe
	TLMessageActionPaymentSent
	TLMessageActionPhoneCall
	TLMessageActionScreenshotTaken
	TLDialog
	TLPhotoEmpty
	TLPhoto
	TLPhotoSizeEmpty
	TLPhotoSize
	TLPhotoCachedSize
	TLGeoPointEmpty
	TLGeoPoint
	TLAuthCheckedPhone
	TLAuthSentCode
	TLAuthAuthorization
	TLAuthExportedAuthorization
	TLInputNotifyPeer
	TLInputNotifyUsers
	TLInputNotifyChats
	TLInputNotifyAll
	TLInputPeerNotifyEventsEmpty
	TLInputPeerNotifyEventsAll
	TLInputPeerNotifySettings
	TLPeerNotifyEventsEmpty
	TLPeerNotifyEventsAll
	TLPeerNotifySettingsEmpty
	TLPeerNotifySettings
	TLPeerSettings
	TLWallPaper
	TLWallPaperSolid
	TLInputReportReasonSpam
	TLInputReportReasonViolence
	TLInputReportReasonPornography
	TLInputReportReasonOther
	TLUserFull
	TLContact
	TLImportedContact
	TLContactBlocked
	TLContactStatus
	TLContactsLink
	TLContactsContactsNotModified
	TLContactsContacts
	TLContactsImportedContacts
	TLContactsBlocked
	TLContactsBlockedSlice
	TLMessagesDialogs
	TLMessagesDialogsSlice
	TLMessagesMessages
	TLMessagesMessagesSlice
	TLMessagesChannelMessages
	TLMessagesChats
	TLMessagesChatsSlice
	TLMessagesChatFull
	TLMessagesAffectedHistory
	TLInputMessagesFilterEmpty
	TLInputMessagesFilterPhotos
	TLInputMessagesFilterVideo
	TLInputMessagesFilterPhotoVideo
	TLInputMessagesFilterPhotoVideoDocuments
	TLInputMessagesFilterDocument
	TLInputMessagesFilterUrl
	TLInputMessagesFilterGif
	TLInputMessagesFilterVoice
	TLInputMessagesFilterMusic
	TLInputMessagesFilterChatPhotos
	TLInputMessagesFilterPhoneCalls
	TLInputMessagesFilterRoundVoice
	TLInputMessagesFilterRoundVideo
	TLInputMessagesFilterMyMentions
	TLUpdateNewMessage
	TLUpdateMessageID
	TLUpdateDeleteMessages
	TLUpdateUserTyping
	TLUpdateChatUserTyping
	TLUpdateChatParticipants
	TLUpdateUserStatus
	TLUpdateUserName
	TLUpdateUserPhoto
	TLUpdateContactRegistered
	TLUpdateContactLink
	TLUpdateNewEncryptedMessage
	TLUpdateEncryptedChatTyping
	TLUpdateEncryption
	TLUpdateEncryptedMessagesRead
	TLUpdateChatParticipantAdd
	TLUpdateChatParticipantDelete
	TLUpdateDcOptions
	TLUpdateUserBlocked
	TLUpdateNotifySettings
	TLUpdateServiceNotification
	TLUpdatePrivacy
	TLUpdateUserPhone
	TLUpdateReadHistoryInbox
	TLUpdateReadHistoryOutbox
	TLUpdateWebPage
	TLUpdateReadMessagesContents
	TLUpdateChannelTooLong
	TLUpdateChannel
	TLUpdateNewChannelMessage
	TLUpdateReadChannelInbox
	TLUpdateDeleteChannelMessages
	TLUpdateChannelMessageViews
	TLUpdateChatAdmins
	TLUpdateChatParticipantAdmin
	TLUpdateNewStickerSet
	TLUpdateStickerSetsOrder
	TLUpdateStickerSets
	TLUpdateSavedGifs
	TLUpdateBotInlineQuery
	TLUpdateBotInlineSend
	TLUpdateEditChannelMessage
	TLUpdateChannelPinnedMessage
	TLUpdateBotCallbackQuery
	TLUpdateEditMessage
	TLUpdateInlineBotCallbackQuery
	TLUpdateReadChannelOutbox
	TLUpdateDraftMessage
	TLUpdateReadFeaturedStickers
	TLUpdateRecentStickers
	TLUpdateConfig
	TLUpdatePtsChanged
	TLUpdateChannelWebPage
	TLUpdateDialogPinned
	TLUpdatePinnedDialogs
	TLUpdateBotWebhookJSON
	TLUpdateBotWebhookJSONQuery
	TLUpdateBotShippingQuery
	TLUpdateBotPrecheckoutQuery
	TLUpdatePhoneCall
	TLUpdateLangPackTooLong
	TLUpdateLangPack
	TLUpdateFavedStickers
	TLUpdateChannelReadMessagesContents
	TLUpdateContactsReset
	TLUpdatesState
	TLUpdatesDifferenceEmpty
	TLUpdatesDifference
	TLUpdatesDifferenceSlice
	TLUpdatesDifferenceTooLong
	TLUpdatesTooLong
	TLUpdateShortMessage
	TLUpdateShortChatMessage
	TLUpdateShort
	TLUpdatesCombined
	TLUpdates
	TLUpdateShortSentMessage
	TLPhotosPhotos
	TLPhotosPhotosSlice
	TLPhotosPhoto
	TLUploadFile
	TLUploadFileCdnRedirect
	TLDcOption
	TLConfig
	TLNearestDc
	TLHelpAppUpdate
	TLHelpNoAppUpdate
	TLHelpInviteText
	TLEncryptedChatEmpty
	TLEncryptedChatWaiting
	TLEncryptedChatRequested
	TLEncryptedChat
	TLEncryptedChatDiscarded
	TLInputEncryptedChat
	TLEncryptedFileEmpty
	TLEncryptedFile
	TLInputEncryptedFileEmpty
	TLInputEncryptedFileUploaded
	TLInputEncryptedFile
	TLInputEncryptedFileBigUploaded
	TLEncryptedMessage
	TLEncryptedMessageService
	TLMessagesDhConfigNotModified
	TLMessagesDhConfig
	TLMessagesSentEncryptedMessage
	TLMessagesSentEncryptedFile
	TLInputDocumentEmpty
	TLInputDocument
	TLDocumentEmpty
	TLDocument
	TLHelpSupport
	TLNotifyPeer
	TLNotifyUsers
	TLNotifyChats
	TLNotifyAll
	TLSendMessageTypingAction
	TLSendMessageCancelAction
	TLSendMessageRecordVideoAction
	TLSendMessageUploadVideoAction
	TLSendMessageRecordAudioAction
	TLSendMessageUploadAudioAction
	TLSendMessageUploadPhotoAction
	TLSendMessageUploadDocumentAction
	TLSendMessageGeoLocationAction
	TLSendMessageChooseContactAction
	TLSendMessageGamePlayAction
	TLSendMessageRecordRoundAction
	TLSendMessageUploadRoundAction
	TLContactsFound
	TLInputPrivacyKeyStatusTimestamp
	TLInputPrivacyKeyChatInvite
	TLInputPrivacyKeyPhoneCall
	TLPrivacyKeyStatusTimestamp
	TLPrivacyKeyChatInvite
	TLPrivacyKeyPhoneCall
	TLInputPrivacyValueAllowContacts
	TLInputPrivacyValueAllowAll
	TLInputPrivacyValueAllowUsers
	TLInputPrivacyValueDisallowContacts
	TLInputPrivacyValueDisallowAll
	TLInputPrivacyValueDisallowUsers
	TLPrivacyValueAllowContacts
	TLPrivacyValueAllowAll
	TLPrivacyValueAllowUsers
	TLPrivacyValueDisallowContacts
	TLPrivacyValueDisallowAll
	TLPrivacyValueDisallowUsers
	TLAccountPrivacyRules
	TLAccountDaysTTL
	TLDocumentAttributeImageSize
	TLDocumentAttributeAnimated
	TLDocumentAttributeSticker
	TLDocumentAttributeVideo
	TLDocumentAttributeAudio
	TLDocumentAttributeFilename
	TLDocumentAttributeHasStickers
	TLMessagesStickersNotModified
	TLMessagesStickers
	TLStickerPack
	TLMessagesAllStickersNotModified
	TLMessagesAllStickers
	TLDisabledFeature
	TLMessagesAffectedMessages
	TLContactLinkUnknown
	TLContactLinkNone
	TLContactLinkHasPhone
	TLContactLinkContact
	TLWebPageEmpty
	TLWebPagePending
	TLWebPage
	TLWebPageNotModified
	TLAuthorization
	TLAccountAuthorizations
	TLAccountNoPassword
	TLAccountPassword
	TLAccountPasswordSettings
	TLAccountPasswordInputSettings
	TLAuthPasswordRecovery
	TLReceivedNotifyMessage
	TLChatInviteEmpty
	TLChatInviteExported
	TLChatInviteAlready
	TLChatInvite
	TLInputStickerSetEmpty
	TLInputStickerSetID
	TLInputStickerSetShortName
	TLStickerSet
	TLMessagesStickerSet
	TLBotCommand
	TLBotInfo
	TLKeyboardButton
	TLKeyboardButtonUrl
	TLKeyboardButtonCallback
	TLKeyboardButtonRequestPhone
	TLKeyboardButtonRequestGeoLocation
	TLKeyboardButtonSwitchInline
	TLKeyboardButtonGame
	TLKeyboardButtonBuy
	TLKeyboardButtonRow
	TLReplyKeyboardHide
	TLReplyKeyboardForceReply
	TLReplyKeyboardMarkup
	TLReplyInlineMarkup
	TLMessageEntityUnknown
	TLMessageEntityMention
	TLMessageEntityHashtag
	TLMessageEntityBotCommand
	TLMessageEntityUrl
	TLMessageEntityEmail
	TLMessageEntityBold
	TLMessageEntityItalic
	TLMessageEntityCode
	TLMessageEntityPre
	TLMessageEntityTextUrl
	TLMessageEntityMentionName
	TLInputMessageEntityMentionName
	TLInputChannelEmpty
	TLInputChannel
	TLContactsResolvedPeer
	TLMessageRange
	TLUpdatesChannelDifferenceEmpty
	TLUpdatesChannelDifferenceTooLong
	TLUpdatesChannelDifference
	TLChannelMessagesFilterEmpty
	TLChannelMessagesFilter
	TLChannelParticipant
	TLChannelParticipantSelf
	TLChannelParticipantCreator
	TLChannelParticipantAdmin
	TLChannelParticipantBanned
	TLChannelParticipantsRecent
	TLChannelParticipantsAdmins
	TLChannelParticipantsKicked
	TLChannelParticipantsBots
	TLChannelParticipantsBanned
	TLChannelParticipantsSearch
	TLChannelsChannelParticipants
	TLChannelsChannelParticipant
	TLHelpTermsOfService
	TLFoundGif
	TLFoundGifCached
	TLMessagesFoundGifs
	TLMessagesSavedGifsNotModified
	TLMessagesSavedGifs
	TLInputBotInlineMessageMediaAuto
	TLInputBotInlineMessageText
	TLInputBotInlineMessageMediaGeo
	TLInputBotInlineMessageMediaVenue
	TLInputBotInlineMessageMediaContact
	TLInputBotInlineMessageGame
	TLInputBotInlineResult
	TLInputBotInlineResultPhoto
	TLInputBotInlineResultDocument
	TLInputBotInlineResultGame
	TLBotInlineMessageMediaAuto
	TLBotInlineMessageText
	TLBotInlineMessageMediaGeo
	TLBotInlineMessageMediaVenue
	TLBotInlineMessageMediaContact
	TLBotInlineResult
	TLBotInlineMediaResult
	TLMessagesBotResults
	TLExportedMessageLink
	TLMessageFwdHeader
	TLAuthCodeTypeSms
	TLAuthCodeTypeCall
	TLAuthCodeTypeFlashCall
	TLAuthSentCodeTypeApp
	TLAuthSentCodeTypeSms
	TLAuthSentCodeTypeCall
	TLAuthSentCodeTypeFlashCall
	TLMessagesBotCallbackAnswer
	TLMessagesMessageEditData
	TLInputBotInlineMessageID
	TLInlineBotSwitchPM
	TLMessagesPeerDialogs
	TLTopPeer
	TLTopPeerCategoryBotsPM
	TLTopPeerCategoryBotsInline
	TLTopPeerCategoryCorrespondents
	TLTopPeerCategoryGroups
	TLTopPeerCategoryChannels
	TLTopPeerCategoryPhoneCalls
	TLTopPeerCategoryPeers
	TLContactsTopPeersNotModified
	TLContactsTopPeers
	TLDraftMessageEmpty
	TLDraftMessage
	TLMessagesFeaturedStickersNotModified
	TLMessagesFeaturedStickers
	TLMessagesRecentStickersNotModified
	TLMessagesRecentStickers
	TLMessagesArchivedStickers
	TLMessagesStickerSetInstallResultSuccess
	TLMessagesStickerSetInstallResultArchive
	TLStickerSetCovered
	TLStickerSetMultiCovered
	TLMaskCoords
	TLInputStickeredMediaPhoto
	TLInputStickeredMediaDocument
	TLGame
	TLInputGameID
	TLInputGameShortName
	TLHighScore
	TLMessagesHighScores
	TLTextEmpty
	TLTextPlain
	TLTextBold
	TLTextItalic
	TLTextUnderline
	TLTextStrike
	TLTextFixed
	TLTextUrl
	TLTextEmail
	TLTextConcat
	TLPageBlockUnsupported
	TLPageBlockTitle
	TLPageBlockSubtitle
	TLPageBlockAuthorDate
	TLPageBlockHeader
	TLPageBlockSubheader
	TLPageBlockParagraph
	TLPageBlockPreformatted
	TLPageBlockFooter
	TLPageBlockDivider
	TLPageBlockAnchor
	TLPageBlockList
	TLPageBlockBlockquote
	TLPageBlockPullquote
	TLPageBlockPhoto
	TLPageBlockVideo
	TLPageBlockCover
	TLPageBlockEmbed
	TLPageBlockEmbedPost
	TLPageBlockCollage
	TLPageBlockSlideshow
	TLPageBlockChannel
	TLPageBlockAudio
	TLPagePart
	TLPageFull
	TLPhoneCallDiscardReasonMissed
	TLPhoneCallDiscardReasonDisconnect
	TLPhoneCallDiscardReasonHangup
	TLPhoneCallDiscardReasonBusy
	TLDataJSON
	TLLabeledPrice
	TLInvoice
	TLPaymentCharge
	TLPostAddress
	TLPaymentRequestedInfo
	TLPaymentSavedCredentialsCard
	TLWebDocument
	TLInputWebDocument
	TLInputWebFileLocation
	TLUploadWebFile
	TLPaymentsPaymentForm
	TLPaymentsValidatedRequestedInfo
	TLPaymentsPaymentResult
	TLPaymentsPaymentVerficationNeeded
	TLPaymentsPaymentReceipt
	TLPaymentsSavedInfo
	TLInputPaymentCredentialsSaved
	TLInputPaymentCredentials
	TLAccountTmpPassword
	TLShippingOption
	TLInputStickerSetItem
	TLInputPhoneCall
	TLPhoneCallEmpty
	TLPhoneCallWaiting
	TLPhoneCallRequested
	TLPhoneCallAccepted
	TLPhoneCall
	TLPhoneCallDiscarded
	TLPhoneConnection
	TLPhoneCallProtocol
	TLPhonePhoneCall
	TLUploadCdnFileReuploadNeeded
	TLUploadCdnFile
	TLCdnPublicKey
	TLCdnConfig
	TLLangPackString
	TLLangPackStringPluralized
	TLLangPackStringDeleted
	TLLangPackDifference
	TLLangPackLanguage
	TLChannelAdminRights
	TLChannelBannedRights
	TLChannelAdminLogEventActionChangeTitle
	TLChannelAdminLogEventActionChangeAbout
	TLChannelAdminLogEventActionChangeUsername
	TLChannelAdminLogEventActionChangePhoto
	TLChannelAdminLogEventActionToggleInvites
	TLChannelAdminLogEventActionToggleSignatures
	TLChannelAdminLogEventActionUpdatePinned
	TLChannelAdminLogEventActionEditMessage
	TLChannelAdminLogEventActionDeleteMessage
	TLChannelAdminLogEventActionParticipantJoin
	TLChannelAdminLogEventActionParticipantLeave
	TLChannelAdminLogEventActionParticipantInvite
	TLChannelAdminLogEventActionParticipantToggleBan
	TLChannelAdminLogEventActionParticipantToggleAdmin
	TLChannelAdminLogEventActionChangeStickerSet
	TLChannelAdminLogEvent
	TLChannelsAdminLogResults
	TLChannelAdminLogEventsFilter
	TLPopularContact
	TLCdnFileHash
	TLMessagesFavedStickersNotModified
	TLMessagesFavedStickers
	TLInvokeAfterMsg
	TLInvokeAfterMsgs
	TLInitConnection
	TLInvokeWithLayer
	TLInvokeWithoutUpdates
	TLAuthCheckPhone
	TLAuthSendCode
	TLAuthResendCode
	TLAccountSendChangePhoneCode
	TLAccountSendConfirmPhoneCode
	TLAuthSignUp
	TLAuthSignIn
	TLAuthImportAuthorization
	TLAuthImportBotAuthorization
	TLAuthCheckPassword
	TLAuthRecoverPassword
	TLAuthLogOut
	TLAuthResetAuthorizations
	TLAuthSendInvites
	TLAuthBindTempAuthKey
	TLAuthCancelCode
	TLAuthDropTempAuthKeys
	TLAccountRegisterDevice
	TLAccountUnregisterDevice
	TLAccountUpdateNotifySettings
	TLAccountResetNotifySettings
	TLAccountUpdateStatus
	TLAccountReportPeer
	TLAccountCheckUsername
	TLAccountDeleteAccount
	TLAccountSetAccountTTL
	TLAccountUpdateDeviceLocked
	TLAccountResetAuthorization
	TLAccountUpdatePasswordSettings
	TLAccountConfirmPhone
	TLContactsDeleteContacts
	TLContactsBlock
	TLContactsUnblock
	TLContactsResetTopPeerRating
	TLContactsResetSaved
	TLMessagesSetTyping
	TLMessagesReportSpam
	TLMessagesHideReportSpam
	TLMessagesDiscardEncryption
	TLMessagesSetEncryptedTyping
	TLMessagesReadEncryptedHistory
	TLMessagesReportEncryptedSpam
	TLMessagesUninstallStickerSet
	TLMessagesEditChatAdmin
	TLMessagesReorderStickerSets
	TLMessagesSaveGif
	TLMessagesSetInlineBotResults
	TLMessagesEditInlineBotMessage
	TLMessagesSetBotCallbackAnswer
	TLMessagesSaveDraft
	TLMessagesReadFeaturedStickers
	TLMessagesSaveRecentSticker
	TLMessagesClearRecentStickers
	TLMessagesSetInlineGameScore
	TLMessagesToggleDialogPin
	TLMessagesReorderPinnedDialogs
	TLMessagesSetBotShippingResults
	TLMessagesSetBotPrecheckoutResults
	TLMessagesFaveSticker
	TLUploadSaveFilePart
	TLUploadSaveBigFilePart
	TLHelpSaveAppLog
	TLHelpSetBotUpdatesStatus
	TLChannelsReadHistory
	TLChannelsReportSpam
	TLChannelsEditAbout
	TLChannelsCheckUsername
	TLChannelsUpdateUsername
	TLChannelsSetStickers
	TLChannelsReadMessageContents
	TLBotsAnswerWebhookJSONQuery
	TLPaymentsClearSavedInfo
	TLPhoneReceivedCall
	TLPhoneSaveCallDebug
	TLAuthExportAuthorization
	TLAuthRequestPasswordRecovery
	TLAccountGetNotifySettings
	TLAccountUpdateProfile
	TLAccountUpdateUsername
	TLAccountChangePhone
	TLContactsImportCard
	TLAccountGetWallPapers
	TLAccountGetPrivacy
	TLAccountSetPrivacy
	TLAccountGetAccountTTL
	TLAccountGetAuthorizations
	TLAccountGetPassword
	TLAccountGetPasswordSettings
	TLAccountGetTmpPassword
	TLUsersGetUsers
	TLUsersGetFullUser
	TLContactsGetStatuses
	TLContactsGetContacts
	TLContactsImportContacts
	TLContactsDeleteContact
	TLContactsGetBlocked
	TLContactsExportCard
	TLMessagesGetMessagesViews
	TLContactsSearch
	TLContactsResolveUsername
	TLContactsGetTopPeers
	TLMessagesGetMessages
	TLMessagesGetHistory
	TLMessagesSearch
	TLMessagesSearchGlobal
	TLMessagesGetUnreadMentions
	TLChannelsGetMessages
	TLMessagesGetDialogs
	TLMessagesReadHistory
	TLMessagesDeleteMessages
	TLMessagesReadMessageContents
	TLChannelsDeleteMessages
	TLMessagesDeleteHistory
	TLChannelsDeleteUserHistory
	TLMessagesReceivedMessages
	TLMessagesSendMessage
	TLMessagesSendMedia
	TLMessagesForwardMessages
	TLMessagesEditChatTitle
	TLMessagesEditChatPhoto
	TLMessagesAddChatUser
	TLMessagesDeleteChatUser
	TLMessagesCreateChat
	TLMessagesForwardMessage
	TLMessagesImportChatInvite
	TLMessagesStartBot
	TLMessagesToggleChatAdmins
	TLMessagesMigrateChat
	TLMessagesSendInlineBotResult
	TLMessagesEditMessage
	TLMessagesGetAllDrafts
	TLMessagesSetGameScore
	TLMessagesSendScreenshotNotification
	TLHelpGetAppChangelog
	TLChannelsCreateChannel
	TLChannelsEditAdmin
	TLChannelsEditTitle
	TLChannelsEditPhoto
	TLChannelsJoinChannel
	TLChannelsLeaveChannel
	TLChannelsInviteToChannel
	TLChannelsDeleteChannel
	TLChannelsToggleInvites
	TLChannelsToggleSignatures
	TLChannelsUpdatePinnedMessage
	TLChannelsEditBanned
	TLPhoneDiscardCall
	TLPhoneSetCallRating
	TLMessagesGetPeerSettings
	TLMessagesGetChats
	TLMessagesGetCommonChats
	TLMessagesGetAllChats
	TLChannelsGetChannels
	TLChannelsGetAdminedPublicChannels
	TLMessagesGetFullChat
	TLChannelsGetFullChannel
	TLMessagesGetDhConfig
	TLMessagesRequestEncryption
	TLMessagesAcceptEncryption
	TLMessagesSendEncrypted
	TLMessagesSendEncryptedFile
	TLMessagesSendEncryptedService
	TLMessagesReceivedQueue
	TLPhotosDeletePhotos
	TLMessagesGetAllStickers
	TLMessagesGetMaskStickers
	TLMessagesGetWebPagePreview
	TLMessagesUploadMedia
	TLMessagesExportChatInvite
	TLChannelsExportInvite
	TLMessagesCheckChatInvite
	TLMessagesGetStickerSet
	TLStickersCreateStickerSet
	TLStickersRemoveStickerFromSet
	TLStickersChangeStickerPosition
	TLStickersAddStickerToSet
	TLMessagesInstallStickerSet
	TLMessagesGetDocumentByHash
	TLMessagesSearchGifs
	TLMessagesGetSavedGifs
	TLMessagesGetInlineBotResults
	TLMessagesGetMessageEditData
	TLMessagesGetBotCallbackAnswer
	TLMessagesGetPeerDialogs
	TLMessagesGetPinnedDialogs
	TLMessagesGetFeaturedStickers
	TLMessagesGetRecentStickers
	TLMessagesGetArchivedStickers
	TLMessagesGetAttachedStickers
	TLMessagesGetGameHighScores
	TLMessagesGetInlineGameHighScores
	TLMessagesGetWebPage
	TLMessagesGetFavedStickers
	TLUpdatesGetState
	TLUpdatesGetDifference
	TLUpdatesGetChannelDifference
	TLPhotosUpdateProfilePhoto
	TLPhotosUploadProfilePhoto
	TLPhotosGetUserPhotos
	TLUploadGetFile
	TLUploadGetWebFile
	TLUploadGetCdnFile
	TLUploadReuploadCdnFile
	TLUploadGetCdnFileHashes
	TLHelpGetConfig
	TLHelpGetNearestDc
	TLHelpGetAppUpdate
	TLHelpGetInviteText
	TLHelpGetSupport
	TLHelpGetTermsOfService
	TLHelpGetCdnConfig
	TLChannelsGetParticipants
	TLChannelsGetParticipant
	TLChannelsExportMessageLink
	TLChannelsGetAdminLog
	TLBotsSendCustomRequest
	TLPhoneGetCallConfig
	TLPaymentsGetPaymentForm
	TLPaymentsGetPaymentReceipt
	TLPaymentsValidateRequestedInfo
	TLPaymentsSendPaymentForm
	TLPaymentsGetSavedInfo
	TLPhoneRequestCall
	TLPhoneAcceptCall
	TLPhoneConfirmCall
	TLLangpackGetLangPack
	TLLangpackGetDifference
	TLLangpackGetStrings
	TLLangpackGetLanguages
	MsgsAck
	BadMsgNotification
	MsgsStateReq
	MsgsStateInfo
	MsgsAllInfo
	MsgDetailedInfo
	MsgResendReq
	RpcError
	RpcDropAnswer
	FutureSalt
	FutureSalts
	Pong
	DestroySessionRes
	NewSession
	HttpWait
	IpPort
	Help_ConfigSimple
	TLMsgsAck
	TLBadMsgNotification
	TLBadServerSalt
	TLMsgsStateReq
	TLMsgsStateInfo
	TLMsgsAllInfo
	TLMsgDetailedInfo
	TLMsgNewDetailedInfo
	TLMsgResendReq
	TLRpcError
	TLRpcAnswerUnknown
	TLRpcAnswerDroppedRunning
	TLRpcAnswerDropped
	TLFutureSalt
	TLFutureSalts
	TLPong
	TLDestroySessionOk
	TLDestroySessionNone
	TLNewSessionCreated
	TLHttpWait
	TLIpPort
	TLHelpConfigSimple
	TLRpcDropAnswer
	TLGetFutureSalts
	TLPing
	TLPingDelayDisconnect
	TLDestroySession
	TLContestSaveDeveloperInfo
*/
package mtproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Type forward declarations
type Bool struct {
	// Types that are valid to be assigned to Payload:
	//	*Bool_BoolFalse
	//	*Bool_BoolTrue
	Payload isBool_Payload `protobuf_oneof:"payload"`
}

func (m *Bool) Reset()                    { *m = Bool{} }
func (m *Bool) String() string            { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()               {}
func (*Bool) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isBool_Payload interface {
	isBool_Payload()
}

type Bool_BoolFalse struct {
	BoolFalse *TLBoolFalse `protobuf:"bytes,1,opt,name=boolFalse,oneof"`
}
type Bool_BoolTrue struct {
	BoolTrue *TLBoolTrue `protobuf:"bytes,2,opt,name=boolTrue,oneof"`
}

func (*Bool_BoolFalse) isBool_Payload() {}
func (*Bool_BoolTrue) isBool_Payload()  {}

func (m *Bool) GetPayload() isBool_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Bool) GetBoolFalse() *TLBoolFalse {
	if x, ok := m.GetPayload().(*Bool_BoolFalse); ok {
		return x.BoolFalse
	}
	return nil
}

func (m *Bool) GetBoolTrue() *TLBoolTrue {
	if x, ok := m.GetPayload().(*Bool_BoolTrue); ok {
		return x.BoolTrue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Bool) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Bool_OneofMarshaler, _Bool_OneofUnmarshaler, _Bool_OneofSizer, []interface{}{
		(*Bool_BoolFalse)(nil),
		(*Bool_BoolTrue)(nil),
	}
}

func _Bool_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Bool)
	// payload
	switch x := m.Payload.(type) {
	case *Bool_BoolFalse:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BoolFalse); err != nil {
			return err
		}
	case *Bool_BoolTrue:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BoolTrue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Bool.Payload has unexpected type %T", x)
	}
	return nil
}

func _Bool_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Bool)
	switch tag {
	case 1: // payload.boolFalse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TLBoolFalse)
		err := b.DecodeMessage(msg)
		m.Payload = &Bool_BoolFalse{msg}
		return true, err
	case 2: // payload.boolTrue
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TLBoolTrue)
		err := b.DecodeMessage(msg)
		m.Payload = &Bool_BoolTrue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Bool_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Bool)
	// payload
	switch x := m.Payload.(type) {
	case *Bool_BoolFalse:
		s := proto.Size(x.BoolFalse)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Bool_BoolTrue:
		s := proto.Size(x.BoolTrue)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type True struct {
	// Types that are valid to be assigned to Payload:
	//	*True_True
	Payload isTrue_Payload `protobuf_oneof:"payload"`
}

func (m *True) Reset()                    { *m = True{} }
func (m *True) String() string            { return proto.CompactTextString(m) }
func (*True) ProtoMessage()               {}
func (*True) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isTrue_Payload interface {
	isTrue_Payload()
}

type True_True struct {
	True *TLTrue `protobuf:"bytes,1,opt,name=true,oneof"`
}

func (*True_True) isTrue_Payload() {}

func (m *True) GetPayload() isTrue_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *True) GetTrue() *TLTrue {
	if x, ok := m.GetPayload().(*True_True); ok {
		return x.True
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*True) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _True_OneofMarshaler, _True_OneofUnmarshaler, _True_OneofSizer, []interface{}{
		(*True_True)(nil),
	}
}

func _True_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*True)
	// payload
	switch x := m.Payload.(type) {
	case *True_True:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.True); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("True.Payload has unexpected type %T", x)
	}
	return nil
}

func _True_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*True)
	switch tag {
	case 1: // payload.true
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TLTrue)
		err := b.DecodeMessage(msg)
		m.Payload = &True_True{msg}
		return true, err
	default:
		return false, nil
	}
}

func _True_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*True)
	// payload
	switch x := m.Payload.(type) {
	case *True_True:
		s := proto.Size(x.True)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Error struct {
	// Types that are valid to be assigned to Payload:
	//	*Error_Error
	Payload isError_Payload `protobuf_oneof:"payload"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isError_Payload interface {
	isError_Payload()
}

type Error_Error struct {
	Error *TLError `protobuf:"bytes,1,opt,name=error,oneof"`
}

func (*Error_Error) isError_Payload() {}

func (m *Error) GetPayload() isError_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Error) GetError() *TLError {
	if x, ok := m.GetPayload().(*Error_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Error) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Error_OneofMarshaler, _Error_OneofUnmarshaler, _Error_OneofSizer, []interface{}{
		(*Error_Error)(nil),
	}
}

func _Error_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Error)
	// payload
	switch x := m.Payload.(type) {
	case *Error_Error:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Error.Payload has unexpected type %T", x)
	}
	return nil
}

func _Error_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Error)
	switch tag {
	case 1: // payload.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TLError)
		err := b.DecodeMessage(msg)
		m.Payload = &Error_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Error_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Error)
	// payload
	switch x := m.Payload.(type) {
	case *Error_Error:
		s := proto.Size(x.Error)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Null struct {
	// Types that are valid to be assigned to Payload:
	//	*Null_Null
	Payload isNull_Payload `protobuf_oneof:"payload"`
}

func (m *Null) Reset()                    { *m = Null{} }
func (m *Null) String() string            { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()               {}
func (*Null) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isNull_Payload interface {
	isNull_Payload()
}

type Null_Null struct {
	Null *TLNull `protobuf:"bytes,1,opt,name=null,oneof"`
}

func (*Null_Null) isNull_Payload() {}

func (m *Null) GetPayload() isNull_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Null) GetNull() *TLNull {
	if x, ok := m.GetPayload().(*Null_Null); ok {
		return x.Null
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Null) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Null_OneofMarshaler, _Null_OneofUnmarshaler, _Null_OneofSizer, []interface{}{
		(*Null_Null)(nil),
	}
}

func _Null_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Null)
	// payload
	switch x := m.Payload.(type) {
	case *Null_Null:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Null); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Null.Payload has unexpected type %T", x)
	}
	return nil
}

func _Null_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Null)
	switch tag {
	case 1: // payload.null
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TLNull)
		err := b.DecodeMessage(msg)
		m.Payload = &Null_Null{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Null_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Null)
	// payload
	switch x := m.Payload.(type) {
	case *Null_Null:
		s := proto.Size(x.Null)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// boolFalse#bc799737 = Bool;
type TLBoolFalse struct {
}

func (m *TLBoolFalse) Reset()                    { *m = TLBoolFalse{} }
func (m *TLBoolFalse) String() string            { return proto.CompactTextString(m) }
func (*TLBoolFalse) ProtoMessage()               {}
func (*TLBoolFalse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// boolTrue#997275b5 = Bool;
type TLBoolTrue struct {
}

func (m *TLBoolTrue) Reset()                    { *m = TLBoolTrue{} }
func (m *TLBoolTrue) String() string            { return proto.CompactTextString(m) }
func (*TLBoolTrue) ProtoMessage()               {}
func (*TLBoolTrue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// true#3fedd339 = True;
type TLTrue struct {
}

func (m *TLTrue) Reset()                    { *m = TLTrue{} }
func (m *TLTrue) String() string            { return proto.CompactTextString(m) }
func (*TLTrue) ProtoMessage()               {}
func (*TLTrue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// error#c4b9f9bb code:int text:string = Error;
type TLError struct {
	Code int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *TLError) Reset()                    { *m = TLError{} }
func (m *TLError) String() string            { return proto.CompactTextString(m) }
func (*TLError) ProtoMessage()               {}
func (*TLError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TLError) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TLError) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// null#56730bcc = Null;
type TLNull struct {
}

func (m *TLNull) Reset()                    { *m = TLNull{} }
func (m *TLNull) String() string            { return proto.CompactTextString(m) }
func (*TLNull) ProtoMessage()               {}
func (*TLNull) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*Bool)(nil), "mtproto.Bool")
	proto.RegisterType((*True)(nil), "mtproto.True")
	proto.RegisterType((*Error)(nil), "mtproto.Error")
	proto.RegisterType((*Null)(nil), "mtproto.Null")
	proto.RegisterType((*TLBoolFalse)(nil), "mtproto.TL_boolFalse")
	proto.RegisterType((*TLBoolTrue)(nil), "mtproto.TL_boolTrue")
	proto.RegisterType((*TLTrue)(nil), "mtproto.TL_true")
	proto.RegisterType((*TLError)(nil), "mtproto.TL_error")
	proto.RegisterType((*TLNull)(nil), "mtproto.TL_null")
}

func init() { proto.RegisterFile("schema.tl.core_types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xbb, 0xd1, 0xda, 0xf5, 0x4d, 0x45, 0x83, 0xc2, 0xd0, 0x8b, 0xf4, 0x20, 0xf3, 0x92,
	0x43, 0xc5, 0x83, 0x07, 0x2f, 0x05, 0xa5, 0x07, 0x15, 0x29, 0xbd, 0x8f, 0xb4, 0x0b, 0x3a, 0x48,
	0xfb, 0x4a, 0x9a, 0x82, 0x03, 0xff, 0x78, 0x79, 0x69, 0x37, 0xc3, 0xf0, 0xd4, 0xaf, 0xef, 0xfb,
	0x7e, 0xf9, 0x5e, 0x02, 0x57, 0x5d, 0xf5, 0x25, 0x6b, 0xc1, 0x8d, 0xe2, 0x15, 0x6a, 0xb9, 0x32,
	0xdb, 0x56, 0x76, 0xbc, 0xd5, 0x68, 0x90, 0x85, 0xb5, 0xb1, 0x22, 0xfe, 0x01, 0x3f, 0x45, 0x54,
	0xec, 0x01, 0xa2, 0x12, 0x51, 0xbd, 0x08, 0xd5, 0xc9, 0xc5, 0xe4, 0x66, 0xb2, 0x9c, 0x27, 0x97,
	0x7c, 0x0c, 0xf1, 0xe2, 0x75, 0xb5, 0x37, 0x33, 0x2f, 0xff, 0x4b, 0xb2, 0x04, 0x66, 0xf4, 0x53,
	0xe8, 0x5e, 0x2e, 0xa6, 0x96, 0xba, 0x38, 0xa4, 0xc8, 0xcb, 0xbc, 0x7c, 0x9f, 0x4b, 0x23, 0x08,
	0x5b, 0xb1, 0x55, 0x28, 0xd6, 0xf1, 0x23, 0xf8, 0x34, 0x62, 0xb7, 0xe0, 0x1b, 0x3a, 0x62, 0x28,
	0x3e, 0x73, 0x8f, 0x30, 0x03, 0x6e, 0x7d, 0x17, 0x7d, 0x82, 0xe0, 0x59, 0x6b, 0xd4, 0xec, 0x0e,
	0x02, 0x49, 0x62, 0x84, 0xcf, 0x5d, 0xd8, 0x1a, 0x99, 0x97, 0x0f, 0x89, 0x83, 0xe6, 0xf7, 0x5e,
	0x29, 0x6a, 0x6e, 0x7a, 0xa5, 0xfe, 0x6b, 0xa6, 0x39, 0x35, 0xd3, 0xd7, 0x45, 0x4f, 0xe1, 0xd8,
	0x7d, 0x90, 0xf8, 0x04, 0xe6, 0xce, 0x55, 0xe3, 0x08, 0xc2, 0x71, 0xed, 0x38, 0x81, 0xd9, 0x6e,
	0x09, 0xc6, 0xc0, 0xaf, 0x70, 0x3d, 0x5c, 0x31, 0xc8, 0xad, 0xa6, 0x99, 0x91, 0xdf, 0xc6, 0xbe,
	0x5c, 0x94, 0x5b, 0x3d, 0xe2, 0xb6, 0x73, 0x09, 0xd7, 0x15, 0xd6, 0xbc, 0x91, 0x65, 0xaf, 0xc4,
	0xa6, 0xe6, 0xb2, 0xf9, 0xdc, 0x34, 0x72, 0xb7, 0x62, 0x1a, 0xbe, 0x15, 0x1f, 0x24, 0xb2, 0x69,
	0x79, 0x64, 0x27, 0xf7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x92, 0xa7, 0x7d, 0xf3, 0x01,
	0x00, 0x00,
}
