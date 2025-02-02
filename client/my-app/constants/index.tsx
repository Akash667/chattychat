import exp from "constants";
import { env } from "process";

const API_URL =  env.API_URL ||  'http://localhost:3000/';


export { API_URL }
