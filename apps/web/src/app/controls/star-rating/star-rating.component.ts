import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-star-rating',
  templateUrl: './star-rating.component.html',
  styleUrls: ['./star-rating.component.scss'],
})
export class StarRatingComponent implements OnInit {
  //
  @Input() stars = 0;
  @Input() reviews = 0;
  starsClass: string[] = [];
  //
  constructor() {}

  ngOnInit(): void {
    for (let i = 0; i < 5; i++) {
      if (i < this.stars) this.starsClass.push('star-on');
      else this.starsClass.push('star-off');
    }
  }
}
