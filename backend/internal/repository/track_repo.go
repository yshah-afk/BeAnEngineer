package repository

import (
	"context"

	"github.com/mastery-hub/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TrackRepo struct {
	coll *mongo.Collection
}

func NewTrackRepo(db *mongo.Database) *TrackRepo {
	return &TrackRepo{coll: db.Collection("tracks")}
}

func (r *TrackRepo) FindAll(ctx context.Context) ([]models.Track, error) {
	opts := options.Find().SetSort(bson.D{{Key: "order", Value: 1}})
	cursor, err := r.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tracks []models.Track
	if err := cursor.All(ctx, &tracks); err != nil {
		return nil, err
	}
	return tracks, nil
}

func (r *TrackRepo) FindBySlug(ctx context.Context, slug string) (*models.Track, error) {
	var track models.Track
	err := r.coll.FindOne(ctx, bson.M{"slug": slug}).Decode(&track)
	if err != nil {
		return nil, err
	}
	return &track, nil
}

func (r *TrackRepo) FindLessonBySlug(ctx context.Context, trackSlug, moduleSlug, lessonSlug string) (*models.Track, *models.Module, *models.Lesson, error) {
	track, err := r.FindBySlug(ctx, trackSlug)
	if err != nil {
		return nil, nil, nil, err
	}

	for i := range track.Modules {
		mod := &track.Modules[i]
		if mod.Slug != moduleSlug {
			continue
		}
		for j := range mod.Lessons {
			les := &mod.Lessons[j]
			if les.Slug == lessonSlug {
				return track, mod, les, nil
			}
		}
	}
	return nil, nil, nil, mongo.ErrNoDocuments
}

func (r *TrackRepo) FindLessonBySlugFlat(ctx context.Context, lessonSlug string) (*models.Track, *models.Module, *models.Lesson, error) {
	filter := bson.M{"modules.lessons.slug": lessonSlug}
	var track models.Track
	err := r.coll.FindOne(ctx, filter).Decode(&track)
	if err != nil {
		return nil, nil, nil, err
	}
	for i := range track.Modules {
		mod := &track.Modules[i]
		for j := range mod.Lessons {
			les := &mod.Lessons[j]
			if les.Slug == lessonSlug {
				return &track, mod, les, nil
			}
		}
	}
	return nil, nil, nil, mongo.ErrNoDocuments
}

func (r *TrackRepo) Upsert(ctx context.Context, track *models.Track) error {
	opts := options.Replace().SetUpsert(true)
	_, err := r.coll.ReplaceOne(ctx, bson.M{"slug": track.Slug}, track, opts)
	return err
}

func (r *TrackRepo) UpdateLesson(ctx context.Context, trackSlug, moduleSlug string, lesson models.Lesson) error {
	filter := bson.M{
		"slug":                 trackSlug,
		"modules.slug":        moduleSlug,
		"modules.lessons.slug": lesson.Slug,
	}
	update := bson.M{
		"$set": bson.M{
			"modules.$[mod].lessons.$[les]": lesson,
		},
	}
	opts := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"mod.slug": moduleSlug},
			bson.M{"les.slug": lesson.Slug},
		},
	})
	_, err := r.coll.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *TrackRepo) AddLesson(ctx context.Context, trackSlug, moduleSlug string, lesson models.Lesson) error {
	filter := bson.M{
		"slug":          trackSlug,
		"modules.slug":  moduleSlug,
	}
	update := bson.M{
		"$push": bson.M{
			"modules.$.lessons": lesson,
		},
	}
	_, err := r.coll.UpdateOne(ctx, filter, update)
	return err
}

func (r *TrackRepo) RemoveLesson(ctx context.Context, trackSlug, moduleSlug, lessonSlug string) error {
	filter := bson.M{
		"slug":         trackSlug,
		"modules.slug": moduleSlug,
	}
	update := bson.M{
		"$pull": bson.M{
			"modules.$.lessons": bson.M{"slug": lessonSlug},
		},
	}
	_, err := r.coll.UpdateOne(ctx, filter, update)
	return err
}
