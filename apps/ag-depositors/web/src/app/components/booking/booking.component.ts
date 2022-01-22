import { Component, Input, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';

@Component({
   selector: 'app-booking',
   templateUrl: './booking.component.html',
   styleUrls: ['./booking.component.scss'],
})
export class BookingComponent implements OnInit {
   @Input() item: any;
   constructor(public activeModal: NgbActiveModal) {}

   ngOnInit(): void {}

   book() {}
}
