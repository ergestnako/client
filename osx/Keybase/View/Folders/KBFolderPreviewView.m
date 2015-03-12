//
//  KBFolderPreviewView.m
//  Keybase
//
//  Created by Gabriel on 3/11/15.
//  Copyright (c) 2015 Gabriel Handford. All rights reserved.
//

#import "KBFolderPreviewView.h"

@interface KBFolderPreviewView ()
@property KBImageView *imageView;
@property KBLabel *nameLabel;
@property KBButton *openButton;
@end

@implementation KBFolderPreviewView

- (void)viewInit {
  [super viewInit];

  _imageView = [[KBImageView alloc] init];
  [self addSubview:_imageView];

  _nameLabel = [[KBLabel alloc] init];
  [self addSubview:_nameLabel];

  _openButton = [KBButton buttonWithText:@"Open in Finder" style:KBButtonStyleToolbar];
  [self addSubview:_openButton];

  YOSelf yself = self;
  self.viewLayout = [YOLayout layoutWithLayoutBlock:^(id<YOLayout> layout, CGSize size) {
    CGFloat x = 6;
    CGFloat y = 2;

    y += [layout centerWithSize:CGSizeMake(120, 120) frame:CGRectMake(0, y, size.width, 0) view:yself.imageView].size.height + 8;

    y += [layout sizeToFitVerticalInFrame:CGRectMake(x, y, size.width - x, 20) view:yself.nameLabel].size.height + 20;

    y += [layout centerWithSize:CGSizeMake(200, 0) frame:CGRectMake(40, y, size.width - 80, 0) view:yself.openButton].size.height + 30;

    return CGSizeMake(size.width, y);
  }];
}

- (void)setFolder:(KBFolder *)folder {
  NSImage *image = KBImageForFolder(folder);
  image.size = CGSizeMake(120, 120);
  self.imageView.image = image;
  [self.nameLabel setText:folder.name style:KBLabelStyleHeader alignment:NSCenterTextAlignment lineBreakMode:NSLineBreakByCharWrapping];
  [self setNeedsLayout];
}

@end
