import { Injectable } from '@angular/core';
import { WarehouseDetail, IDetailService } from '../interfaces/detail';

@Injectable({
  providedIn: 'root',
})
export class MDetailService implements IDetailService {
  //
  private item!: WarehouseDetail;

  constructor() {
    this.initTestData();
  }

  async get(id: string): Promise<WarehouseDetail> {
    return this.item;
  }

  private initTestData() {
    this.item = {
      propertyID: 'ONESHEDS0003',
      noOfViews: 204,
      available: true,
      images: [
        'assets/mock/wh1.jpeg',
        'assets/mock/wh2.jpeg',
        'assets/mock/wh3.jpeg',
        'assets/mock/wh4.jpeg',
        'assets/mock/wh5.jpeg',
      ],

      name: 'Sai Vinayak Warehousing Ltd',
      address:
        '18, 3rd Cross Road, NS Palya, BTM Layout Stage 2, Bengaluru, 560076, Karnataka, India',
      localtion: { latitude: 37.4232, longitude: -122.0853 },
      description: `Industrial shed available for lease on prime location good connectivity
                from the main Dadari road. The shed is located on wide road, Its good for
                all industrial use like logistic, and other industry. The surrounded area
                is occupied with all MNC like TCS, Samsung mother diary etc. Interested
                client can contact us for more details and option all across Noida.`,
      rating: 4,
      reviewsCount: 12,

      sharable: true,
      insured: true,
      greenWarehouse: true,
      warehouseType: 'General warehousing', // General warehousing, Dedicated, Custom Bonded, Management warehousing, Cold Storage, Build to Make

      cost: {
        rentalValue: 12,
        depositAmount: 130000,
        minimumTenure: 2,
        bookingAmount: 50000,
        price: 14, //- per sq ft
        rentAppreciation: 8, //Per Year(%)
      },

      facilityCharges: {
        electricityCharges: 10,
        waterCharges: 2,
        cleaningCharges: 1,
        otherCharges: 2,
      },

      //sqft
      size: {
        height: 25,
        sideheight: 20,
        centerHeight: 25,
        plinthHeight: 20,
        superArea: 6000,
        plotArea: 6200,
        coveredArea: 6000,
      },

      power: {
        phases: '3 Phase',
        backupPower: '3 gensets',
        consumption: 'not known',
      },

      finedDetails: {
        flooringType: 'Tiles',
        floor: 2,
        passageAdequacy: 'Good',
        numberofGates: 2,
        cornerPlot: true,
        onRoad: false,
        gateSize: 30, //sqft
        warehouseDirection: 'North-South',
        facing: 'North',
        furnishedStatus: 'Furnished',
      },

      registeration: {
        registrationOffice: 'Hoskote Industrial Area',
        WDRARegistration: 'WDRA00012',
        registrationDate: '12/04/2008',
        registrationNumber: 'REG000023',
        PAN: 'PANKNNK88U',
        TAN: 'TANKNNK88U',
        TIN: 'TINKNNK88U',
        GSTIN: 'GSTINKNNK88U',
        CSTIN: 'CGSTINNNK88U',
        CIN: 'CINNNK88U',
      },

      accrediatation: [
        { item: 'AEO', value: true },
        { item: 'BIS/CWC/FCI', value: false },
        { item: 'C-PAT', value: true },
        { item: 'FSSAI License', value: true },
        { item: 'ISO 14001', value: false },
        { item: 'ISO 9001', value: false },
        { item: 'MCX Approval', value: false },
        { item: 'Fire NOC', value: true },
        {
          item: 'NOC from the local authority ',
          value: true,
        },
        { item: 'NOC from Emergency Services', value: true },

        { item: 'NOC from the State Pollution Control Board', value: true },
        { item: 'WDRA Registration', value: false },
      ],

      surroundings: {
        shoppingMall: ['GGN Central Mall', 'Spencer Ciry Mall'],
        commercialArea: ['Hoskote Industrial Area'],
        ResidentialArea: ['Chanasandara', 'Kadugodi'],
        landmarks: ['Amazon warehouse Hoskote'],
      },

      distanceFrom: {
        important: {
          policeStation: 6,
          fireStation: 10,
          weighbridge: 2,
          highway: 3,
        },

        hospitals: ['City Hospital Hoskote', 'Navjeevan Chaildcare'],
        metro: ['Whitefield metro station'],
        railway: ['Whitefield railway station'],
        busStop: ['Kadugodi'],
        taxiStation: [],
      },

      amenities: [
        { item: 'Weighbridge', value: true },
        { item: '24x7 Security', value: true },
        { item: 'CCTV camera Monitored', value: true },
        { item: 'Inventory Management System', value: true },
        { item: 'Logistics / Transportation', value: true },
        { item: 'Power Backup', value: true },
        { item: 'Rainwater Harvesting', value: false },
        { item: 'Reserved Parking', value: true },
        { item: 'Water Storage', value: true },
        { item: 'Private Terrace/Garden', value: true },
        { item: 'Vaastu Compliant', value: true },
        { item: 'Visitor Parking', value: false },
        { item: 'Maintenance Staff', value: true },
        { item: 'RO Water System', value: false },
        { item: 'Toilet', value: true },
        { item: '24 Hour Access', value: true },
        { item: 'Admin Block', value: true },
        { item: 'Boundary Wall', value: false },
        { item: 'Canteen', value: true },
        { item: 'Electrical Connection', value: true },
        { item: 'Fire Protection', value: true },
        { item: 'Fork Lift', value: true },
        { item: 'Guard Room', value: true },
        { item: 'Loading Dock', value: true },
        { item: 'Local Labor availability / accommodation', value: false },
        { item: 'Racking', value: true },
        { item: 'Secure Facility', value: true },
        { item: 'Water Connection', value: true },
        { item: 'Drivers Lounge', value: true },
        { item: 'Solar Panel Capable', value: true },
        { item: 'Wifi Connectivity', value: true },
        { item: 'Dock', value: true },
      ],

      dates: {
        commencedOn: '12/04/1998',
        availableFrom: '12/04/1998',
        listedOn: '12/04/2022',
        bookedTill: 'NA',
        refurbished: 'Yes',
        ageofConstruction: 24,
      },

      reviews: [
        {
          companyName: 'Jayalcove Pharma',
          userName: 'Mohan Lal',
          rating: 4,
          feedback:
            'All the facilities maintioned online are available in very good condition',
          date: '12/04/2022',
        },
        {
          companyName: 'Neelkamal Furniture House',
          userName: 'Sohan Lal',
          rating: 5,
          feedback:
            'All the facilities maintioned online are available in very good condition',
          date: '16/04/2021',
        },
      ],

      equipments: [
        { item: 'Forklift', value: true },
        { item: 'Hydra', value: true },
        { item: 'Pallet Jacks', value: true },
        { item: 'Integrated Dock Levelers', value: true },
        { item: 'Edge of Dock Levelers', value: true },
        { item: 'Truck Restraints', value: true },
        { item: 'Dock Seals and Shelters', value: true },
        { item: 'Dock Boards and Plates', value: true },
        { item: 'Wheel Chocks', value: true },
        { item: 'Bumpers', value: true },
        { item: 'Lights', value: true },
        { item: 'Strip Doors and Air Curtains', value: true },
        { item: 'Large Ceiling Fans', value: true },
        { item: 'Yard Ramps', value: true },
        { item: 'Cranes and Hoists', value: true },
        { item: 'Dollies', value: true },
        { item: 'Work Benches', value: true },
        { item: 'Utility Carts', value: true },
        { item: 'Trucks', value: true },
        { item: 'Casters', value: true },
        { item: 'Totes and Bins', value: true },
      ],

      scales: [
        { item: 'Floor Scales', value: true },
        { item: 'Small Parts Scale', value: true },
        { item: 'High-Speed Check-weighing Conveyor Scale', value: true },
        { item: 'Pallet Scale', value: true },
        { item: 'Check-weighing and Cubing', value: true },
      ],

      safetyEquipments: [
        { item: 'Emergency Wash Station', value: true },
        { item: 'Antifatique Mats', value: true },
        { item: 'Barrier Rails', value: true },
        { item: 'Bollards', value: true },
        { item: 'Column Protectors', value: true },
        { item: 'Wire Partitions', value: true },
        { item: 'Traffic Visibility Mirrors', value: true },
        { item: 'Handrails', value: true },
        { item: 'Miscellaneous Equipment', value: true },
      ],

      storageOptions: [
        { item: 'On ground', value: true },
        { item: 'Racking', value: true },
        { item: 'Secured Room(s)', value: true },
        { item: 'Pick Module(s)', value: true },
        { item: 'Dedicated Room(s)', value: true },
        { item: 'Single Stacking', value: true },
        { item: 'Double Stacking', value: false },
        { item: 'Pallet less Stacking', value: true },
        { item: 'Raking System', value: true },
        { item: 'Slotted Angle Racks', value: false },
        { item: 'Bin Storage', value: false },
        { item: 'Palletize Ground Storage', value: true },
        { item: 'Cold Storage', value: false },
      ],

      vas: [
        { item: 'Pallet Storage', value: true },
        { item: 'Carton Pick', value: true },
        { item: 'Pallet Rebuild', value: true },
        { item: 'Cross Dock', value: true },
        { item: 'Last Mile Delivery', value: true },
        { item: 'Bar Coding', value: true },
        { item: 'Labelling', value: true },
        { item: 'Kitting', value: true },
        { item: 'Return Logistics', value: true },
        { item: 'Local Distribution', value: true },
        { item: 'Freight Forwarding', value: true },
        { item: 'Repacking Services', value: true },
        { item: 'Intra/Inter -City Transport', value: true },
      ],

      industrieServed: [
        { item: 'Agriculture & Prepared Products', value: true },
        { item: 'Apparel, Footwear & Textiles', value: true },
        { item: 'Automotive & Aerospace', value: true },
        { item: 'Base Metals', value: true },
        { item: 'Consumer Products and Mass Merchandising', value: true },
        { item: 'Electronics', value: true },
        { item: 'Industrial & Manufacturing Materials', value: true },
        { item: 'Machinery', value: true },
        { item: 'Petroleum, Natural Gas & Minerals', value: true },
        { item: 'Pharmaceutical, Health and Chemicals', value: true },
      ],

      prohibbited: ['Acids', 'Explosives'],

      documents: {
        //Before Start of Construction
        before: [
          { item: 'Non Agriculture Approval', value: 'doc0001' },
          { item: 'Land Conversion Approval', value: 'doc0001' },
          { item: 'Environmental Approval', value: 'doc0001' },
          { item: 'Master Plan Approval', value: 'doc0001' },
          { item: 'Building Plan Approval', value: 'doc0001' },
          { item: 'Preliminary Fire NOC', value: 'doc0001' },
          { item: 'Consent to Establish', value: 'doc0001' },
          { item: 'Central Ground Water Board Clearance', value: 'doc0001' },
          {
            item: 'Contract Labor Regulation and Abolition Act',
            value: 'doc0001',
          },
          {
            item: 'Approval from Chief Electrical Inspector to Government',
            value: 'doc0001',
          },
          {
            item: 'Access Approval (If Applicable National Highway, State Highway, Private Body)',
            value: 'doc0001',
          },
        ],
        // After Construction/Normal
        after: [
          { item: 'Final Fire NOC', value: 'doc0001' },
          { item: 'OC/BCC - Occupational Certificate', value: 'doc0001' },
          { item: 'Consent to Operate (CTO)', value: 'doc0001' },
          {
            item: 'Approval for Commencement of Operations from local Panchayat (if applicable)',
            value: 'doc0001',
          },
          { item: 'Master Plan of Area', value: 'doc0001' },
          { item: 'Floor Plan', value: 'doc0001' },
          { item: 'Building Plan', value: 'doc0001' },
          { item: 'Warehouse Registration Certificate', value: 'doc0001' },
        ],
      },

      bookingsHistory: [],

      companyDetails: {
        name: 'Saivinayak Warehouses Ltd',
        url: 'www.saivinayakwh.com',
        mobile: '22979879798',
        email: 'contact@saivinayakwh.com',
        chat: 'saivinayakwh@skype',
      },

      contactDetails: {
        name: 'Manhardas',
        role: 'Manager',
        mobile: '98884784783',
        email: 'manohardas@saivinayakwh.com',
        chat: 'NA',
      },

      localityRatings: {
        enviromnent: {
          neighborhood: 4,
          roads: 2,
          safety: 3,
          cleanliness: 4,
        },
        commuting: {
          publicTransport: 4,
          parking: 5,
          connectivity: 4,
          traffic: 3,
        },
        placesOfInterest: {
          schools: 4,
          restaurants: 3,
          hospital: 4,
          market: 2,
        },
      },
    };
  }
}
