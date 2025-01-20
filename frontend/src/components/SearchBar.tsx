import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";

interface SearchBarProps {
    onSearch: (type: 'id' | 'city', value: string) => void;
}

export function SearchBar({ onSearch }: SearchBarProps) {
    const [searchType, setSearchType] = useState<'id' | 'city'>('id');
    const [searchValue, setSearchValue] = useState('');

    const handleSearch = () => {
        if (searchValue.trim()) {
            onSearch(searchType, searchValue.trim());
        }
    };

    return (
        <div className="flex gap-2 mb-4">
            <Select
                value={searchType}
                onValueChange={(value: 'id' | 'city') => setSearchType(value)}
            >
                <SelectTrigger className="w-[180px]">
                    <SelectValue placeholder="Filtrar por" />
                </SelectTrigger>
                <SelectContent>
                    <SelectItem value="id">ID</SelectItem>
                    <SelectItem value="city">Cidade</SelectItem>
                </SelectContent>
            </Select>
            <Input
                placeholder="Buscar..."
                value={searchValue}
                onChange={(e) => setSearchValue(e.target.value)}
                className="max-w-sm"
            />
            <Button onClick={handleSearch}>Buscar</Button>
        </div>
    );
}