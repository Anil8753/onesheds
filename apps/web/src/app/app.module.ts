import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgxSliderModule } from '@angular-slider/ngx-slider';

import { HttpClientModule } from '@angular/common/http';
import { MaterialModule } from './material.modules';
import { NgxGalleryModule } from 'ngx-gallery-9';

import { ToolbarComponent } from './components/toolbar/toolbar.component';
import { SearchCtrlComponent } from './components/search-ctrl/search-ctrl.component';
import { GooglePlaceModule } from 'ngx-google-places-autocomplete';
import { StatsComponent } from './components/stats/stats.component';
import { OurServicesComponent } from './components/our-services/our-services.component';
import { SocialMediaComponent } from './components/social-media/social-media.component';
import { FooterComponent } from './components/footer/footer.component';
import { ListingComponent } from './components/listing/listing.component';
import { LandingComponent } from './components/landing/landing.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { FilterComponent } from './components/listing/filter/filter.component';
import { RangeSliderComponent } from './controls/range-slider/range-slider.component';
import { RatingsComponent } from './components/listing/filter/ratings/ratings.component';
import { TypesComponent } from './components/listing/filter/types/types.component';
import { ListItemComponent } from './components/listing/list-item/list-item.component';
import { StarRatingComponent } from './controls/star-rating/star-rating.component';
import { SearchResultsHeaderComponent } from './components/listing/search-results-header/search-results-header.component';
import { DetailsComponent } from './components/details/details.component';
import { VerificationComponent } from './controls/verification/verification.component';
import { ChipsListComponent } from './controls/chips-list/chips-list.component';
import { YesNoPipe } from './pipes/yes-no.pipe';
import { DetailsTopComponent } from './components/details/details-top/details-top.component';
import { GalleryComponent } from './components/details/gallery/gallery.component';
import { TextFieldComponent } from './controls/text-field/text-field.component';
import { DetailsMidComponent } from './components/details/details-mid/details-mid.component';
import { SliceStringPipe } from './pipes/slice-string.pipe';
import { TextListComponent } from './controls/text-list/text-list.component';
import { DetailsMidActionsComponent } from './components/details/details-mid-actions/details-mid-actions.component';
import { DetailsBriefComponent } from './components/details/details-brief/details-brief.component';
import { DetailsSurrondingComponent } from './components/details/details-surronding/details-surronding.component';
import { DetailsDistanceComponent } from './components/details/details-distance/details-distance.component';
import { DetailsLocalityRatingComponent } from './components/details/details-locality-rating/details-locality-rating.component';
import { UserReviewComponent } from './controls/user-review/user-review.component';
import { DetailsContactComponent } from './components/details/details-contact/details-contact.component';
// import { GoogleMapsModule } from '@angular/google-maps';

@NgModule({
  declarations: [
    AppComponent,
    ToolbarComponent,
    SearchCtrlComponent,
    StatsComponent,
    OurServicesComponent,
    SocialMediaComponent,
    FooterComponent,
    ListingComponent,
    LandingComponent,
    PageNotFoundComponent,
    FilterComponent,
    RangeSliderComponent,
    RatingsComponent,
    TypesComponent,
    ListItemComponent,
    StarRatingComponent,
    SearchResultsHeaderComponent,
    DetailsComponent,
    VerificationComponent,
    ChipsListComponent,
    YesNoPipe,
    DetailsTopComponent,
    GalleryComponent,
    TextFieldComponent,
    DetailsMidComponent,
    SliceStringPipe,
    TextListComponent,
    DetailsMidActionsComponent,
    DetailsBriefComponent,
    DetailsSurrondingComponent,
    DetailsDistanceComponent,
    DetailsLocalityRatingComponent,
    UserReviewComponent,
    DetailsContactComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MaterialModule,
    FormsModule,

    ReactiveFormsModule,

    HttpClientModule,
    NgxSliderModule,
    NgxGalleryModule,
    GooglePlaceModule,
    // GoogleMapsModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
