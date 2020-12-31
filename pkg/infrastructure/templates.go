package infrastructure

const baseIndex = `export * from './__name__Criteria';
export * from './__name__Model';
export * from './__name__Repository';
`

func baseController() string {
	return `import * as express from 'express';
import { __name__ } from '../../Domain';
import { Add__name__, Get__name__, Remove__name__, Update__name__ } from '../../Application';
import { __name__Criteria, __name__Repository } from '../persistence/__repo__';

const router = express.Router();
const __var__Route = '/__route__';

router.get(__var__Route, async (req: express.Request, res: express.Response) => {
    try {
        const get__name__ = new Get__name__(new __name__Repository(), new __name__Criteria(req.query));
        const __var__ = await get__name__.dispatch();

        res.json(__var__);
    } catch (error) {
        res.status(error.statusCode || 503).json({ error: error.message || error });
    }
});

router.post(__var__Route, async (req: express.Request, res: express.Response) => {
    try {
        const add__name__ = new Add__name__(
            new __name__Repository(),
            new __name__(req.body),
        );
        await add__name__.dispatch();

        res.status(201).json({ message: '__name__ create success' });
    } catch (error) {
        res.status(error.statusCode || 503).json({ error: error.message || error });
    }
});

router.put(__var__Route, async (req: express.Request, res: express.Response) => {
    try {
        const update__name__ = new Update__name__(
            new __name__Repository(),
            new __name__(req.body),
        );
        await update__name__.dispatch();

        res.json({ message: '__name__ update success' });
    } catch (error) {
        res.status(error.statusCode || 503).json({ error: error.message || error });
    }
});

router.delete(__var__Route, async (req: express.Request, res: express.Response) => {
    try {
        const remove__name__ = new Remove__name__(new __name__Repository(), req.query.id);
        await remove__name__.dispatch();

        res.json({ message: '__name__ remove success' });
    } catch (error) {
        res.status(error.statusCode || 503).json({ error: error.message || error });
    }
});

export { router as __name__Router };
`
}
