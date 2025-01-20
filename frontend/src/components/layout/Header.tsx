import { MdLocationOn } from "react-icons/md";

export function Header() {
    return (
        <header className="border-b">
            <div className="container mx-auto px-4 py-4">
                <div className="flex items-center gap-2">
                    <MdLocationOn className="w-6 h-6" />
                    <h1 className="text-xl font-bold">Reverse Geocoding</h1>
                </div>
            </div>
        </header>
    );
}