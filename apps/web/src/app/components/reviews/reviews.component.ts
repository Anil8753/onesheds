import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-reviews',
  templateUrl: './reviews.component.html',
  styleUrls: ['./reviews.component.scss'],
})
export class ReviewsComponent implements OnInit {
  @Input()
  data!: {
    total: number;
    list: {
      companyName: string;
      userName: string;
      rating: number;
      feedback: string;
      date: string;
    }[];
  };
  constructor() {}

  ngOnInit(): void {}
}