import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-details-contact',
  templateUrl: './details-contact.component.html',
  styleUrls: ['./details-contact.component.scss'],
})
export class DetailsContactComponent implements OnInit {
  //
  @Input() data:
    | {
        name: string;
        role: string;
        mobile: string;
        email: string;
        chat: string;
      }
    | undefined;

  constructor() {}

  ngOnInit(): void {}
}
