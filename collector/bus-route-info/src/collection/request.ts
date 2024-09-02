import { parseStringPromise } from "xml2js";

export const requestAndParseXml = async <T> (url: string): Promise<T> => {
    try {
        const response = await fetch(url);
        if (!response.ok) throw new Error(`Error occurred duriuing request, status: ${response.status}`);
        const xmlText = await response.text();
        const result = await parseStringPromise(xmlText, {explicitArray: false});
        return result as T;
    } catch (error) {
        throw new Error(`Error request or parsing XML: ${error.message}`);
    }
}
