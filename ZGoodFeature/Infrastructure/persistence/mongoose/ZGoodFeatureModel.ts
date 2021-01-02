import * as mongoose from "mongoose";
import { IZGoodFeature } from "../../../Domain";

interface ZGoodFeatureDoc extends mongoose.Document, IZGoodFeature {
	id: string;
	username: string;
	group_id: string;
}

type ZGoodFeatureModel = mongoose.Model<ZGoodFeatureDoc>

const zGoodFeatureSchema = new mongoose.Schema<IZGoodFeature>(
	{
		id_user: {
			type: String,
			required: true
		},
		username: {
			type: String,
			required: true
		},
		group_id: {
			type: String,
			required: true
		},
	  lastname: {
			type: String,
			required: true
		},
	  age: {
			type: Number,
			required: true
		},
	  active: {
			type: Boolean,
			required: true
		},
	  firstname: {
			type: String,
			required: true
		},
  },
	{
		toJSON: {
			transform(doc, ret) {
				delete ret.__v;
				ret.id = ret._id;
			}
		}
	}
);

const ZGoodFeatureModel = mongoose.model<ZGoodFeatureDoc, ZGoodFeatureModel>("zGoodFeature", zGoodFeatureSchema);

export { ZGoodFeatureModel };
