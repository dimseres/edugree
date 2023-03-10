import {PrismaClient} from "@prisma/client"

export interface ISeeder {
    seed(client: PrismaClient): Function|Function[],
}