import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class WarehouseDataService {
  items: WarehouseItem[] = [];
  constructor() {
    this.initTestData();
  }

  async getWareHouses(params: any): Promise<WarehouseItem[]> {
    return this.items;
  }

  private initTestData() {
    this.items = [
      {
        name: 'Navata Warehouses',
        address:
          '18, 3rd Cross Road, NS Palya, BTM Layout Stage 2, Bengaluru, 560076, Karnataka, India',
        description:
          'Cold storage warehouses enable prescription and over-the-counter medications',
        thumbUrl:
          'https://navata.com/cms/wp-content/uploads/2021/08/warehousing.jpg',
        usersRating: 5,
        userReviews: 18,
        verified: true,
        securityDeposit: 150000,
        warehouseType: 'Cold Storage',
        rate: 12,
        totalSize: 2000,
        blockedSize: 500,
        minTenure: 3,
        brokerage: 20,
        wdraReg: false,
        ageYear: '2006',
        prohibittedItems: ['Explosives', 'Grains', 'Acids'],
        keyAttractions: [
          'Security camera coverage',
          'Transport facility',
          'NH-44 2 km',
          'BTM Firestation 3 km',
          'Transport city 3 km',
        ],
        hasVideo: true,
      },
      {
        name: 'Goodluck Warehouses',
        address:
          '76, Thavarekere Main Road, BTM Layout Stage 1, Bengaluru, 560029, Karnataka, India',
        description:
          'An essential component of the supply chain are these supply chain facilities',
        thumbUrl:
          'https://navata.com/cms/wp-content/uploads/2021/08/AW-Blog-What-Is-Public-Warehousing.jpg',
        usersRating: 3,
        userReviews: 9,
        verified: true,
        warehouseType: 'Private Warehouses',
        rate: 32,
        securityDeposit: 20000,
        totalSize: 5000,
        blockedSize: 1000,
        minTenure: 1,
        brokerage: 0,
        wdraReg: true,
        ageYear: '1966',
        prohibittedItems: ['Oils', 'Explosives'],
        keyAttractions: [
          'National highway 5 km',
          'Transport city 3 km',
          'WIFI coverage',
          'Weighingbridge',
        ],
        hasVideo: false,
      },
      {
        name: 'Rayman logistics and warehouses',
        address:
          '403-13A, 13th Cross Road, Sri Venkateshwara Layout, Bengaluru, 560068, Karnataka, India',
        description:
          ' lets the government authorities maintain control over private companies',
        thumbUrl:
          'https://navata.com/cms/wp-content/uploads/2021/08/Tins-row-2048x1153.jpg',
        usersRating: 1,
        userReviews: 0,
        verified: false,
        warehouseType: 'Bonded Storage',
        rate: 22,
        securityDeposit: 120000,
        totalSize: 9000,
        blockedSize: 4000,
        minTenure: 5,
        brokerage: 15,
        wdraReg: true,
        ageYear: '1986',
        prohibittedItems: [],
        keyAttractions: ['Fire station 5 km', 'NH-21 6 km', 'KIAB 2 km'],
        hasVideo: true,
      },
    ];
  }

  randomNumber(min: number, max: number) {
    return Math.floor(Math.random() * (max - min) + min);
  }
}

export interface WarehouseItem {
  name: string;
  address: string;
  description: string;
  usersRating: number;
  userReviews: number;
  verified: boolean;
  thumbUrl: string;
  warehouseType: string;
  rate: number;
  securityDeposit: number;
  totalSize: number;
  blockedSize: number;
  minTenure: number;
  brokerage: number;
  wdraReg: boolean;
  ageYear: string;
  prohibittedItems: string[];
  keyAttractions: string[];
  hasVideo: boolean;
}
