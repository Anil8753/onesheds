export interface IItemListService {
  get(location: string, km: number): Promise<ListItem[]>;
}

export interface ListItem {
  id: string;
  name: string;
  address: string;
  description: string;
  usersRating: number;
  userReviews: number;
  localityRating: number;
  verified: boolean;
  thumbUrl: string;
  warehouseType: string;
  rate: number;
  rateNegotiable: boolean;
  securityDeposit: number;
  totalSize: number;
  blockedSize: number;
  minTenure: number;
  wdraReg: boolean;
  operatingSince: string;
  shared: boolean;
  insured: boolean;

  location: { latitude: number; longitude: number };

  greenWarehouse: boolean;
  onRoad: boolean;
  parking: number;

  amenities: string[];
  hasVideo: boolean;
}
