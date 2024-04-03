import { NextPage } from "next";
import dynamic from "next/dynamic";
import "./globals.css";

const AdminApp = dynamic(() => import("@/components/AdminApp"), { ssr: false });

const Home: NextPage = () => <AdminApp />;

export default Home;
