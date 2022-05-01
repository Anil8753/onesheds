import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-details-locality-rating',
  templateUrl: './details-locality-rating.component.html',
  styleUrls: ['./details-locality-rating.component.scss'],
})
export class DetailsLocalityRatingComponent implements OnInit {
  @Input() data:
    | {
        enviromnent: {
          neighborhood: number;
          roads: number;
          safety: number;
          cleanliness: number;
        };
        commuting: {
          publicTransport: number;
          parking: number;
          connectivity: number;
          traffic: number;
        };
        placesOfInterest: {
          schools: number;
          restaurants: number;
          hospital: number;
          market: number;
        };
      }
    | undefined;
  constructor() {}

  ngOnInit(): void {}
}
