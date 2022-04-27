import { Injectable } from '@angular/core';
import { IItemListService, ListItem } from '../interfaces/list';

@Injectable({
  providedIn: 'root',
})
export class MListService implements IItemListService {
  //
  items: ListItem[] = [];

  constructor() {
    this.initTestData();
  }

  async get(locaton: string, km: number): Promise<ListItem[]> {
    return this.items;
  }

  private initTestData() {
    this.items = [
      {
        id: '1236',
        name: 'Navata Warehouses',
        address:
          '18, 3rd Cross Road, NS Palya, BTM Layout Stage 2, Bengaluru, 560076, Karnataka, India',
        description:
          'The property comprises of 2 floors, ground floor and basement, 7500 sqft each with additional provision to make storage / accomodation on the on top of ground floor. Seperate access for basement and provision of goods lift has been provided.',
        thumbUrl:
          'https://navata.com/cms/wp-content/uploads/2021/08/warehousing.jpg',
        usersRating: 5,
        userReviews: 18,
        localityRating: 5,
        verified: true,
        securityDeposit: 150000,
        warehouseType: 'Cold Storage',
        rate: 12,
        rateNegotiable: true,

        totalSize: 2000,
        blockedSize: 500,
        minTenure: 3,
        wdraReg: false,
        operatingSince: '2006',
        shared: false,
        insured: true,

        location: { latitude: 37.6153, longitude: -122.39 },

        greenWarehouse: true,
        onRoad: true,
        parking: 2,

        amenities: [
          'Weighbridge',
          '24x7 Security',
          'CCTV camera Monitored',
          'Inventory Management System',
          'Logistics / Transportation',
          'Admin Block',
          'Canteen',
        ],
        hasVideo: true,
      },
      {
        id: '1235',
        name: 'Goodluck Warehouses',
        address:
          '76, Thavarekere Main Road, BTM Layout Stage 1, Bengaluru, 560029, Karnataka, India',
        description:
          'Commercial space in chandigarh (Ut) for warehouse/ office/ retail. 82 feet front with private parking space. 300 metre from hallomajra chowk and 1 km from tribune chowk. Approx 10,000 square feet of total carpet area in (Basement,ground, first floor and second floor).',
        thumbUrl:
          'https://navata.com/cms/wp-content/uploads/2021/08/AW-Blog-What-Is-Public-Warehousing.jpg',
        usersRating: 3,
        userReviews: 9,
        localityRating: 4,
        verified: true,
        warehouseType: 'Private Warehouses',
        rate: 32,
        rateNegotiable: false,
        securityDeposit: 20000,
        totalSize: 5000,
        blockedSize: 1000,
        minTenure: 1,
        wdraReg: true,
        operatingSince: '1966',
        shared: true,
        insured: true,

        location: { latitude: 37.4289, longitude: -122.1697 },

        greenWarehouse: false,
        onRoad: true,
        parking: 5,

        amenities: [
          'CCTV camera Monitored',
          'Inventory Management System',
          'Logistics / Transportation',
          'Drivers Lounge',
          'Solar Panel Capable',
          'Wifi Connectivity',
        ],
        hasVideo: false,
      },
      {
        id: '1234',
        name: 'Rayman logistics and warehouses',
        address:
          '403-13A, 13th Cross Road, Sri Venkateshwara Layout, Bengaluru, 560068, Karnataka, India',
        description:
          'Ware house available for sell. The property comes with a good construction quality which ages above 10 years. Property is ready to move . Equipped with 1 bathroom. Located in hutta colony. Property is built in 3000 sq.Ft.(Builtup area) . Available at an expected price of "sixty lakh rupees".',
        thumbUrl:
          'https://navata.com/cms/wp-content/uploads/2021/08/Tins-row-2048x1153.jpg',
        usersRating: 1,
        userReviews: 0,
        localityRating: 2,
        verified: false,
        warehouseType: 'Bonded Storage',
        rate: 22,
        rateNegotiable: true,
        securityDeposit: 120000,
        totalSize: 9000,
        blockedSize: 4000,
        minTenure: 5,
        wdraReg: true,
        operatingSince: '1986',
        shared: false,
        insured: true,

        location: { latitude: 37.4232, longitude: -122.0853 },

        greenWarehouse: true,
        onRoad: true,
        parking: 0,

        amenities: [
          'CCTV camera Monitored',
          'Inventory Management System',
          'Logistics / Transportation',
          'Admin Block',
        ],
        hasVideo: true,
      },
    ];
  }
}