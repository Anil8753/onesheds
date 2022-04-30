export interface IDetailService {
  get(id: string): Promise<WarehouseDetail>;
}

export interface WarehouseDetail {
  propertyID: string;
  noOfViews: number;
  available: boolean;
  images: string[];

  name: string;
  address: string;
  localtion: { latitude: number; longitude: number };
  description: string;

  rating: number;
  reviewsCount: number;

  sharable: boolean;
  insured: boolean;
  greenWarehouse: boolean;
  warehouseType: string; // General warehousing, Dedicated, Custom Bonded, Management warehousing, Cold Storage, Build to Make

  cost: {
    rentalValue: number;
    depositAmount: number;
    minimumTenure: number;
    bookingAmount: number;
    price: number; //- per sq ft
    rentAppreciation: number; //Per Year(%)
  };

  facilityCharges: {
    electricityCharges: number;
    waterCharges: number;
    cleaningCharges: number;
    otherCharges: number;
  };

  //sqft
  size: {
    height: number;
    sideheight: number;
    centerHeight: number;
    plinthHeight: number;
    superArea: number;
    plotArea: number;
    coveredArea: number;
  };

  power: {
    phases: string;
    backupPower: string;
    consumption: string;
  };

  finedDetails: {
    flooringType: string;
    floor: number;
    passageAdequacy: string;
    numberofGates: number;
    cornerPlot: boolean;
    onRoad: boolean;
    gateSize: number; //sqft
    warehouseDirection: string;
    facing: string;
    furnishedStatus: string;
  };

  registeration: {
    registrationOffice: string;
    WDRARegistration: string;
    registrationDate: string;
    registrationNumber: string;
    PAN: string;
    TAN: string;
    TIN: string;
    GSTIN: string;
    CSTIN: string;
    CIN: string;
  };

  accrediatation: { item: string; value: string }[];

  surroundings: {
    shoppingMall: string[];
    commercialArea: string[];
    ResidentialArea: string[];
    landmarks: string[];
  };

  distanceFrom: {
    important: {
      policeStation: number;
      fireStation: number;
      weighbridge: number;
      highway: number;
    };

    hospitals: string[];
    metro: string[];
    railway: string[];
    busStop: string[];
    taxiStation: string[];
  };

  amenities: { item: string; value: boolean }[];

  dates: {
    commencedOn: string;
    availableFrom: string;
    listedOn: string;
    bookedTill: string;
    refurbished: string;
    ageofConstruction: number;
  };

  reviews: {
    companyName: string;
    userName: string;
    rating: number;
    feedback: string;
    date: string;
  }[];

  equipments: { item: string; value: boolean }[];

  scales: { item: string; value: boolean }[]; // Warehouse Scales & Check-Weighers

  safetyEquipments: { item: string; value: boolean }[]; // Warehouse Safety Equipment

  storageOptions: { item: string; value: boolean }[];

  vas: { item: string; value: boolean }[]; // value added services

  industrieServed: { item: string; value: boolean }[]; // Industries Served

  prohibbited: string[];

  documents: {
    before: { item: string; value: string }[]; //Before Start of Construction
    after: { item: string; value: string }[]; // After Construction/Normal
  };

  bookingsHistory: string[];

  companyDetails: {
    name: string;
    url: string;
    mobile: string;
    email: string;
    chat: string;
  };

  contactDetails: {
    name: string;
    role: string;
    mobile: string;
    email: string;
    chat: string;
  };

  localityRatings: {
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
  };
}
