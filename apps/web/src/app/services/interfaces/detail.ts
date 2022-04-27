export interface IDetailService {
  get(id: string): Promise<Detail>;
}

export interface Detail {
  id: string;
  name: string;
}
