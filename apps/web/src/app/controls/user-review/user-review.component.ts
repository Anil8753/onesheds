import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-user-review',
  templateUrl: './user-review.component.html',
  styleUrls: ['./user-review.component.scss'],
})
export class UserReviewComponent implements OnInit {
  @Input() reviews = 0;
  constructor() {}

  ngOnInit(): void {}
}
