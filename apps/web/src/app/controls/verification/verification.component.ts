import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-verification',
  templateUrl: './verification.component.html',
  styleUrls: ['./verification.component.scss'],
})
export class VerificationComponent implements OnInit {
  //
  @Input() verified: boolean;
  constructor() {
    this.verified = false;
  }

  ngOnInit(): void {}
}
