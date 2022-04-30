import { Component, OnInit } from '@angular/core';
import {
  NgxGalleryOptions,
  NgxGalleryImage,
  NgxGalleryAnimation,
} from 'ngx-gallery-9';

@Component({
  selector: 'app-gallery',
  templateUrl: './gallery.component.html',
  styleUrls: ['./gallery.component.scss'],
})
export class GalleryComponent implements OnInit {
  //
  galleryOptions: NgxGalleryOptions[] = [];
  galleryImages: NgxGalleryImage[] = [];

  constructor() {}

  ngOnInit(): void {
    this.galleryOptions = [
      {
        width: '350px',
        height: '320px',

        thumbnailsColumns: 4,
        thumbnailsRemainingCount: true,
        thumbnailsArrowsAutoHide: true,

        imageAnimation: NgxGalleryAnimation.Slide,
        imageBullets: true,
        imageArrowsAutoHide: true,

        previewCloseOnClick: true,
        previewCloseOnEsc: true,
        previewKeyboardNavigation: true,
      },
      // max-width 800
      {
        breakpoint: 800,
        width: '100%',
        height: '600px',
        imagePercent: 80,
        thumbnailsPercent: 20,
        thumbnailsMargin: 20,
        thumbnailMargin: 20,
      },
      // max-width 400
      {
        breakpoint: 400,
        preview: false,
      },
    ];

    this.galleryImages = [
      {
        small: 'assets/mock/wh1.jpeg',
        medium: 'assets/mock/wh1.jpeg',
        big: 'assets/mock/wh1.jpeg',
      },
      {
        small: 'assets/mock/wh2.jpeg',
        medium: 'assets/mock/wh2.jpeg',
        big: 'assets/mock/wh2.jpeg',
      },
      {
        small: 'assets/mock/wh3.jpeg',
        medium: 'assets/mock/wh3.jpeg',
        big: 'assets/mock/wh3.jpeg',
      },
      {
        small: 'assets/mock/wh4.jpeg',
        medium: 'assets/mock/wh4.jpeg',
        big: 'assets/mock/wh4.jpeg',
      },
      {
        small: 'assets/mock/wh5.jpeg',
        medium: 'assets/mock/wh5.jpeg',
        big: 'assets/mock/wh5.jpeg',
      },
    ];
  }
}
